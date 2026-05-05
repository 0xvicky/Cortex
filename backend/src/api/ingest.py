from fastapi import APIRouter
from src.models.req import RequestRepo
from src.worker.queue import q
from src.storage.storage import create_job
from src.jobs.ingest_job import run_pipeline
from src.utils.utils import generate_job_id
from datetime import date
import time

import uuid

router = APIRouter(prefix="/ingest", tags=["analysis", "ingest"])


@router.post("/repo")
def ingest_repo(req: RequestRepo):
    repo_url = str(req.repo_url)  # normalize early
    user_id = req.user_id
    today = date.today()
    timestamp = int(time.time())
    job_id = generate_job_id(repo_url=repo_url, date=str(today), time=str(timestamp))
    # job_id = str(uuid.uuid4())
    # user_id = str(uuid.uuid4())
    print(q.connection, q.name)
    create_job(job_id, user_id, repo_url)
    print("here ")
    q.enqueue(
        run_pipeline,
        job_id,
        repo_url,
        user_id,
    )

    return {"status": "ok", "job_id": job_id, "user_id": user_id}
