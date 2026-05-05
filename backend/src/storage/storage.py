from src.models.storage import JobModel
from src.models.chunks import ChunkModel
from typing import Dict
from redis import Redis
import json
from typing import Optional
from src.models.file import FileModel

redis = Redis(host="localhost", port=6379, db=0)


def create_job(job_id, user_id, repo_url):
    job = {
        "job_id": job_id,
        "user_id": user_id,
        "status": "PENDING",
        "repo_url": repo_url,
        "repo_summary": None,
    }
    redis.hset(user_id, job_id, json.dumps(job))
    # redis.rpush(user_id, json.dumps(job))


def update_job(job_id, user_id, key, value):
    data = redis.hget(user_id, job_id)
    if not data:
        return
    job = json.loads(data.decode())
    job[key] = value
    redis.hset(user_id, job_id, json.dumps(job))


def store_repo_summary(job_id: str, repo_summary: dict, repo_url: str, user_id: str):
    data = redis.hget(user_id, job_id)  # read from hash, not plain key
    if not data:
        return

    job = json.loads(data.decode())

    job["repo_summary"] = repo_summary
    job["repo_url"] = repo_url

    redis.hset(user_id, job_id, json.dumps(job))  # write back to hash


def get_job(job_id: str, user_id: str):

    data = redis.hget(user_id, job_id)
    return json.loads(data.decode()) if data else None


def get_all_jobs(user_id: str):
    data = redis.hkeys(user_id)
    if not data:
        return None
    return [job_id.decode() for job_id in data]
