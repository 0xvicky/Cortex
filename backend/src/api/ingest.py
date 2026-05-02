from fastapi import APIRouter, UploadFile, File, Form, HTTPException
from src.models.req import RequestRepo
from src.services.ingest import ingestr_svc
from typing import Optional

router = APIRouter(prefix="/ingest", tags=["analysis, ingest"])


@router.post("/repo")
async def ingest_repo(req: RequestRepo):
    repo_url = req.repo_url
    res = await ingestr_svc(str(repo_url))
    return {"status": "ok", "message": res}
