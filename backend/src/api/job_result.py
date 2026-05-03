from fastapi import APIRouter, UploadFile, File, Form, HTTPException
from src.models.req import RequestRepo
from src.services.ingest import ingestr_svc
from typing import Optional
from src.worker.queue import q
from src.storage.storage import get_job
import uuid

router = APIRouter(prefix="/result", tags=["result"])


@router.get("/{job_id}")
def result(job_id: str):
    print(job_id)
    print("in res")
    job = get_job(job_id)
    if not job:
        return {"error": "Job not found"}
    return {"job": job}
