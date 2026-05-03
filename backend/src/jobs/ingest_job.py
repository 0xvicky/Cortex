# app/jobs/ingest_job.py

import asyncio
from src.services.ingest import ingestr_svc
from src.storage.storage import update_status, store_chunks


def run_pipeline(job_id: str, repo_url: str, user_id: str):
    try:
        print("in pipeline")
        update_status(job_id, "RUNNING")

        # 🔥 run async service inside sync wrapper
        chunks = asyncio.run(ingestr_svc(repo_url))

        store_chunks(job_id, chunks)

        update_status(job_id, "COMPLETED")

    except Exception as e:
        update_status(job_id, "FAILED")
        raise e
