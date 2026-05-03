from src.models.storage import JobModel
from src.models.chunks import ChunkModel
from typing import Dict
from redis import Redis
import json
from typing import Optional

redis = Redis(host="localhost", port=6379, db=0)


def create_job(job_id, user_id, status, chunks):
    redis.set(
        job_id,
        json.dumps(
            {
                "job_id": job_id,
                "user_id": user_id,
                "status": status,
                "nChunks": 0,
                "chunks": None,
            }
        ),
    )


def update_status(job_id, status):
    data = redis.get(job_id)
    if not data:
        return

    job = json.loads(data.decode())  # ✅ FIX
    job["status"] = status
    redis.set(job_id, json.dumps(job))


def store_chunks(job_id: str, chunks: list[ChunkModel]):
    data = redis.get(job_id)
    if not data:
        return

    job = json.loads(data.decode())

    # convert chunks → serializable
    job["chunks"] = [c.model_dump() for c in chunks]
    print(len(chunks))
    job["nChunks"] = len(chunks)
    redis.set(job_id, json.dumps(job))


def get_job(job_id: str):
    data: Optional[bytes] = redis.get(job_id)
    return json.loads(data.decode()) if data else None
