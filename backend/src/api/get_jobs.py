from fastapi import APIRouter
from src.models.req import RequestRepo
from src.worker.rq_worker import q
from src.storage.storage import get_all_jobs
from src.jobs.ingest_job import run_pipeline

import uuid

router = APIRouter(prefix="/jobs", tags=["alljobs"])


@router.get("/{user_id}")
def ingest_repo(user_id: str):
    # job_id = str(uuid.uuid4())
    # user_id = str(uuid.uuid4())
    # print(q.connection, q.name)
    # create_job(job_id, user_id, repo_url)
    res = get_all_jobs(user_id)
    return {"res": res}
