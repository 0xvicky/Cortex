from typing import Optional
from fastapi import UploadFile, HTTPException
from pydantic import HttpUrl
from urllib.parse import urlparse
from github import Github, UnknownObjectException
from dotenv import load_dotenv
import httpx
import os
from pathlib import Path
import uuid

load_dotenv()

g = Github(os.getenv("GITHUB_TOKEN"))

# add github proper repo checks using github package


async def ingestr_svc(repo_url: str):
    parsed = urlparse(repo_url)
    if parsed.netloc != "github.com":
        raise HTTPException(status_code=400, detail="Only GitHub repos allowed")

    async with httpx.AsyncClient() as client:
        res = await client.get(repo_url)
        if res.status_code != 200:
            raise HTTPException(status_code=400, detail="Invalid Github Repository")

    # create a local temp folder with unique name to store clone repo
    PATH = "Z:/Code/Golang/Projects/cortex/temp"
    repo_name = uuid.uuid8()
    # clone that repo into that folder
    print(repo_name)
    repo_path = Path(f"{PATH}/{repo_name}")
    repo_path.mkdir(parents=True, exist_ok=True)
    return {"repoUrl": repo_path}
