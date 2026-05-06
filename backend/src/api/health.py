from fastapi import APIRouter
import os
import re
import uuid
from pathlib import Path
from typing import List, Optional
from urllib.parse import urlparse
from src.models.test import TestRepo
from dotenv import load_dotenv
from fastapi import HTTPException
import base64
import json

from github import Github, UnknownObjectException

g = Github(os.getenv("GITHUB_TOKEN"))

router = APIRouter(prefix="/health", tags=["health"])


@router.get("/")
async def health_check():
    return {"status": "ok", "service": "cortex-backend"}


@router.post("/test")
def generate_job_id(r: TestRepo):
    print("this")
    print(r)
    payload = json.dumps({"repo_url": r.repo_url, "date": r.date})
    print(base64.urlsafe_b64encode(payload.encode()).decode())
    return base64.urlsafe_b64encode(payload.encode()).decode()


# async def validate_github_repo(req: RequestRepo):
#     """Raise HTTPException if the URL is not a valid, accessible GitHub repo."""
#     repo_url = str(req.repo_url)
#     parsed = urlparse(repo_url)
#     if parsed.netloc != "github.com":
#         raise HTTPException(status_code=400, detail="Only GitHub repos are allowed")

#     parts = parsed.path.strip("/").split("/")
#     if len(parts) < 2:
#         raise HTTPException(status_code=400, detail="Invalid GitHub repo URL")

#     owner, repo = parts[0], parts[1]
#     try:
#         g.get_repo(f"{owner}/{repo}")
#         return {"res": "Ok"}


#     except UnknownObjectException:
#         raise HTTPException(
#             status_code=404, detail="GitHub repo not found or is private"
#         )
