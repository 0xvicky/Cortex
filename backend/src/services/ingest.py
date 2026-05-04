import os
import re
import uuid
from pathlib import Path
from typing import List, Optional
from urllib.parse import urlparse

from dotenv import load_dotenv
from fastapi import HTTPException
from git import Repo
from github import Github, UnknownObjectException

from src.constants import constants
from src.models.chunks import ChunkModel
from src.models.file import FileModel
from src.services.llm import llm_query
from src.storage.storage import store_repo_summary
from src.utils.utils import should_process
import json

load_dotenv()

g = Github(os.getenv("GITHUB_TOKEN"))

PATH = "Z:/Code/Golang/Projects/cortex/temp"
SKIP_DIRS = constants.SKIP_DIRS

ALLOWED_EXTENSIONS = {
    ".py",
    ".ts",
    ".js",
    ".go",
    ".rs",
    ".java",
    ".md",
    ".yaml",
    ".yml",
    ".toml",
    ".json",
    ".sql",
    ".php",
}

NAME_PATTERNS = [
    (r"readme", 10),
    (r"main\.", 8),
    (r"app\.", 8),
    (r"application\.", 8),
    (r"server\.", 7),
    (r"index\.", 6),
    (r"schema\.", 6),
    (r"config|settings", 6),
    (r"requirements|pyproject|package\.json|go\.mod|cargo\.toml", 7),
    (r"router|routing", 5),
    (r"model", 5),
    (r"dockerfile", 5),
    (r"__init__\.", 3),
]


# ─────────────────────────────────────────────
#  Validation
# ─────────────────────────────────────────────


def validate_github_repo(repo_url: str):
    """Raise HTTPException if the URL is not a valid, accessible GitHub repo."""
    parsed = urlparse(repo_url)
    if parsed.netloc != "github.com":
        raise HTTPException(status_code=400, detail="Only GitHub repos are allowed")

    parts = parsed.path.strip("/").split("/")
    if len(parts) < 2:
        raise HTTPException(status_code=400, detail="Invalid GitHub repo URL")

    owner, repo = parts[0], parts[1]
    try:
        g.get_repo(f"{owner}/{repo}")
        return (owner, repo)
    except UnknownObjectException:
        raise HTTPException(
            status_code=404, detail="GitHub repo not found or is private"
        )


# ─────────────────────────────────────────────
#  File walking
# ─────────────────────────────────────────────


def get_files(local_path: str) -> List[FileModel]:
    """Walk repo and return all processable files."""
    files: List[FileModel] = []
    file_count = 0

    for root, dirs, filenames in os.walk(local_path):
        dirs[:] = [d for d in dirs if d not in SKIP_DIRS]
        for filename in filenames:
            full_path = os.path.join(root, filename)
            if not should_process(full_path):
                continue
            files.append(FileModel(file_no=file_count, file_path=full_path))
            file_count += 1

    return files


# ─────────────────────────────────────────────
#  Heuristic scorer for summary file selection
# ─────────────────────────────────────────────


def score_file(path: str, size: int) -> int:
    """Score a file based on name patterns, depth, and size."""
    lower = path.lower()
    ext = os.path.splitext(lower)[1]

    if ext not in ALLOWED_EXTENSIONS:
        return 0

    score = 0

    # name pattern score
    for pattern, points in NAME_PATTERNS:
        if re.search(pattern, lower):
            score += points

    # depth score — shallower = more important
    depth = lower.count(os.sep)
    score += max(0, 4 - depth)

    # size score
    if 200 < size < 5_000:
        score += 4
    elif 5_000 <= size < 20_000:
        score += 2
    elif size < 200 or size > 50_000:
        score -= 2

    return score


def get_top_files(local_path: str, top_k: int = 10) -> List[dict]:
    """Return top-k scored files for summary generation."""
    scored = []

    for root, dirs, filenames in os.walk(local_path):
        dirs[:] = [d for d in dirs if d not in SKIP_DIRS]
        for filename in filenames:
            full_path = os.path.join(root, filename)
            rel_path = os.path.relpath(full_path, local_path)
            size = os.path.getsize(full_path)
            s = score_file(rel_path, size)
            if s > 0:
                scored.append({"path": rel_path, "full_path": full_path, "score": s})

    scored.sort(key=lambda x: x["score"], reverse=True)
    return scored[:top_k]


# ─────────────────────────────────────────────
#  Summary  (1 LLM call)
# ─────────────────────────────────────────────


def generate_repo_summary(local_path: str, repo_url: str):
    """Pick top files, extract snippets, call LLM once."""
    top_files = get_top_files(local_path)

    snippets = ""
    total_chars = 0
    CHAR_BUDGET = 4_000

    for f in top_files:
        try:
            content = open(f["full_path"], encoding="utf-8", errors="ignore").read()
        except Exception:
            continue
        snippet = content[:500]
        if total_chars + len(snippet) > CHAR_BUDGET:
            break
        snippets += f"\n\n### {f['path']}\n```\n{snippet}\n```"
        total_chars += len(snippet)

    system_prompt = "You are a senior software engineer. Summarize GitHub repositories concisely and accurately."
    user_prompt = f"""Analyze this repository: {repo_url}

        Based on the following key files, return a JSON object only. No explanation, no markdown, no extra text.

        Return exactly this structure, example given below:
        
        {{
        "project_purpose": "2-3 sentences on what this project does",
        "tech_stack": ["Python", "FastAPI", "Redis"],
        "key_components": [
            {{ "name": "main.py", "description": "Entry point of the application" }},
            {{ "name": "src/routes.py", "description": "Handles all API routes" }}
        ],
        "how_to_run": [
            "pip install -r requirements.txt",
            "uvicorn main:app --reload"
        ],
        "architecture": "Brief 1-2 sentence description of how the system is structured",
        "use_cases": ["Use case 1", "Use case 2"]
        }}

        Files:
        {snippets}

            """
    raw = llm_query(system_prompt, user_prompt)
    try:
        # strip backticks if LLM wraps in ```json ... ```
        clean = (
            raw.strip()
            .removeprefix("```json")
            .removeprefix("```")
            .removesuffix("```")
            .strip()
        )
        return json.loads(clean)
    except json.JSONDecodeError:
        # fallback so pipeline never breaks
        return {
            "project_purpose": raw,
            "tech_stack": [],
            "key_components": [],
            "how_to_run": [],
            "architecture": "",
            "use_cases": [],
        }


# ─────────────────────────────────────────────
#  Chunking  (for RAG / Qdrant)
# ─────────────────────────────────────────────


def chunk_files(
    files: List[FileModel], chunk_size: int = 50, overlap: int = 20
) -> List[ChunkModel]:
    """Sliding window chunker for all repo files."""
    chunks: List[ChunkModel] = []
    step = chunk_size - overlap

    for fl in files:
        try:
            with open(fl.file_path, encoding="utf-8", errors="ignore") as f:
                lines = f.readlines()
        except Exception:
            continue

        chunk_id = 0
        i = 0
        while i < len(lines):
            chunk_lines = lines[i : i + chunk_size]
            if not chunk_lines:
                break
            chunks.append(
                ChunkModel(
                    chunk_no=chunk_id,
                    file_no=fl.file_no,
                    file_path=fl.file_path,
                    content="".join(chunk_lines),
                    start_line=i,
                    end_line=i + len(chunk_lines),
                )
            )
            chunk_id += 1
            i += step

    return chunks


# ─────────────────────────────────────────────
#  Main ingestion entry point
# ─────────────────────────────────────────────


def ingestr_svc(repo_url: str, job_id: str) -> List[ChunkModel]:
    # 1. validate
    # owner, repo = validate_github_repo(repo_url)

    # 2. clone
    local_path = Path(f"{PATH}/{job_id}")
    local_path.mkdir(parents=True, exist_ok=True)
    Repo.clone_from(repo_url, local_path)
    print(f"[ingest] cloned → {local_path}")

    # 3. quick summary — 1 LLM call, runs first
    summary = generate_repo_summary(str(local_path), repo_url)

    store_repo_summary(job_id=job_id, repo_summary=summary)
    print(f"[ingest] summary stored for job {job_id}")

    # 4. RAG chunking — all files, sliding window
    files = get_files(str(local_path))
    chunks = chunk_files(files)

    print(f"[ingest] {len(chunks)} chunks ready for embedding")

    return chunks
