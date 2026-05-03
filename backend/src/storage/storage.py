from src.models.storage import JobModel
from src.models.chunks import ChunkModel
from typing import Dict

jobs: Dict[str, JobModel] = {}


def create_job(job_id, user_id, status, chunks):
    print("in job creation")
    jobs[job_id] = JobModel(
        job_id=job_id, user_id=user_id, status="PENDING", chunks=None
    )


def update_status(job_id: str, status: str):
    print("in update")
    if job_id in jobs:
        jobs[job_id].status = status


def store_chunks(job_id: str, chunks: list[ChunkModel]):
    print("in store")
    if job_id in jobs:
        jobs[job_id].chunks = chunks


def get_job(job_id: str):
    print(jobs.get(job_id))
    return jobs.get(job_id)
