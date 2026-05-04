from src.models.storage import JobModel
from src.models.chunks import ChunkModel
from typing import Dict
from redis import Redis
import json
from typing import Optional
from src.models.file import FileModel

redis = Redis(host="localhost", port=6379, db=0)


def create_job(job_id, user_id):
    job = {
        "job_id": job_id,
        "user_id": user_id,
        "status": "PENDING",
        "repo_summary": None,
    }
    redis.set(job_id, json.dumps(job))


def store_repo_summary(job_id: str, repo_summary: dict):
    data = redis.get(job_id)
    if not data:
        return

    job = json.loads(data.decode())

    job["repo_summary"] = repo_summary

    job["status"] = "COMPLETED"

    redis.set(job_id, json.dumps(job))


def get_job(job_id: str):
    data = redis.get(job_id)
    return json.loads(data.decode()) if data else None
