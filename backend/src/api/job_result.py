from fastapi import APIRouter
from src.storage.storage import get_job

router = APIRouter(prefix="/result", tags=["result"])

@router.get("/{user_id}/{job_id}")
def result(user_id: str, job_id: str):
    job = get_job(job_id, user_id)
    if not job:
        return {"error": "Job not found"}
    return {"job": job}
