from fastapi import APIRouter
from src.worker.rq_worker import q
from src.storage.storage import get_all_jobs
from src.jobs.ingest_job import run_pipeline

router = APIRouter(prefix="/jobs", tags=["alljobs"])

@router.get("/{user_id}")
def ingest_repo(user_id: str):

    res = get_all_jobs(user_id)
    return {"res": res}
