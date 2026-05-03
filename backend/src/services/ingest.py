from typing import Optional, List
from fastapi import UploadFile, HTTPException
from pydantic import HttpUrl
from src.models.file import FileModel
from src.models.chunks import ChunkModel
from urllib.parse import urlparse
from src.constants import constants
from github import Github, UnknownObjectException
from git import Repo
from dotenv import load_dotenv
import httpx
import os
from pathlib import Path
from src.utils.utils import should_process
import uuid

load_dotenv()

g = Github(os.getenv("GITHUB_TOKEN"))
SKIP_DIRS = constants.SKIP_DIRS


def get_files(local_path: str) -> List[FileModel]:
    print("in get files")
    files: List[FileModel] = []
    fileCount: int = 0
    for root, dirs, filenames in os.walk(local_path):
        dirs[:] = [d for d in dirs if d not in SKIP_DIRS]
        for file in filenames:
            full_path = os.path.join(root, file)
            if not should_process(full_path):
                continue
            file_tmp: FileModel = FileModel(file_no=fileCount, file_path=full_path)
            files.append(file_tmp)
            fileCount += 1
    return files


def process_file(ls: List[FileModel]) -> List[ChunkModel]:
    print("in chunking")
    chunks: List[ChunkModel] = []

    chunk_size = 50
    overlap = 20
    step = chunk_size - overlap  # 30

    for fl in ls:
        try:
            with open(fl.file_path, "r", encoding="utf-8", errors="ignore") as f:
                lines = f.readlines()
        except Exception:
            continue

        chunk_id = 0
        i = 0

        while i < len(lines):
            chunk_lines = lines[i : i + chunk_size]

            # skip tiny tail if you want
            if not chunk_lines:
                break

            chunks.append(
                ChunkModel(
                    chunk_no=chunk_id,
                    file_no=fl.file_no,
                    file_path=fl.file_path,
                    content="".join(chunk_lines),  # keep original newlines
                    start_line=i,
                    end_line=i + len(chunk_lines),
                )
            )

            chunk_id += 1
            i += step  # 🔥 sliding window

    return chunks


def ingestr_svc(repo_url: str):
    print("in ingest svc")

    parsed = urlparse(repo_url)
    if parsed.netloc != "github.com":
        raise HTTPException(status_code=400, detail="Only GitHub repos allowed")

    # validate repo (use requests or httpx sync)
    import requests

    res = requests.get(repo_url)
    if res.status_code != 200:
        raise HTTPException(status_code=400, detail="Invalid Github Repository")

    PATH = "Z:/Code/Golang/Projects/cortex/temp"
    repo_name = uuid.uuid4()

    print(repo_name)

    local_path = Path(f"{PATH}/{repo_name}")
    local_path.mkdir(parents=True, exist_ok=True)

    cloned_repo = Repo.clone_from(repo_url, local_path)
    print(cloned_repo)

    files = get_files(str(local_path))
    chunks = process_file(files)

    return chunks


# https://github.com/0xvicky/nebula
# https://github.com/0xvicky/temp-conv-rust


# ===================CHECKS======================================#
# add github proper repo checks using github package
