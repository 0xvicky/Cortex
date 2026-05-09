# benchmark.py
import time
import os
from pathlib import Path
from git import Repo
import tempfile
import shutil

# ---- CONFIG ----
TEST_REPOS = [
    "https://github.com/0xvicky/go-mongodb-integration",  # small
    "https://github.com/tiangolo/fastapi",  # medium
    "https://github.com/django/django",  # large
]


# ---- HELPERS ----
def count_files(path):
    return len([f for f in Path(path).rglob("*") if f.is_file()])


def clone_repo(url):
    path = tempfile.mkdtemp()
    Repo.clone_from(url, path)
    return path


def naive_llm_calls(file_count):
    """Naive = 1 LLM call per file"""
    return file_count


def heuristic_llm_calls():
    """Heuristic = always 1 LLM call"""
    return 1


# ---- BENCHMARK ----
results = []

for repo_url in TEST_REPOS:
    print(f"\n📦 Cloning {repo_url}...")

    start = time.time()
    repo_path = clone_repo(repo_url)
    clone_time = time.time() - start

    file_count = count_files(repo_path)
    naive_calls = naive_llm_calls(file_count)
    heuristic_calls = heuristic_llm_calls()
    reduction = ((naive_calls - heuristic_calls) / naive_calls) * 100

    results.append(
        {
            "repo": repo_url.split("/")[-1],
            "files": file_count,
            "naive_calls": naive_calls,
            "heuristic_calls": heuristic_calls,
            "reduction": round(reduction, 2),
            "clone_time": round(clone_time, 2),
        }
    )

    shutil.rmtree(repo_path, ignore_errors=True)

# ---- RESULTS ----
print("\n" + "=" * 60)
print(f"{'Repo':<30} {'Files':>6} {'Naive':>6} {'Ours':>5} {'Saved':>8}")
print("=" * 60)
for r in results:
    print(
        f"{r['repo']:<30} {r['files']:>6} {r['naive_calls']:>6} {r['heuristic_calls']:>5} {r['reduction']:>7}%"
    )
print("=" * 60)
