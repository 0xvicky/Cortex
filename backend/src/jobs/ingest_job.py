# app/jobs/ingest_job.py

from src.services.ingest import ingestr_svc
from src.storage.vector_db import embed_chunks
from src.storage.storage import update_job


def run_pipeline(job_id: str, repo_url: str, user_id: str):
    try:
        print("in pipeline")

        chunks = ingestr_svc(repo_url, job_id, user_id)  # ✅ no asyncio.run

        # store the chunks in vector db for rag retrieval
        res = embed_chunks(job_id, chunks)

        update_job(job_id=job_id, user_id=user_id, key="status", value=res)
        # update_status(job_id, "COMPLETED")

    except Exception as e:
        update_job(job_id=job_id, user_id=user_id, key="status", value="FAILED")
        raise e
