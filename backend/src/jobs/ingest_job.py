# app/jobs/ingest_job.py

from src.services.ingest import ingestr_svc
from src.storage.vector_db import embed_chunks


def run_pipeline(job_id: str, repo_url: str, user_id: str):
    try:
        print("in pipeline")

        chunks = ingestr_svc(repo_url, job_id, user_id)  # ✅ no asyncio.run

        # store the chunks in vector db for rag retrieval
        embed_chunks(job_id, chunks)
        # update_status(job_id, "COMPLETED")

    except Exception as e:
        # update_status(job_id, "FAILED")
        raise e


# Z:\Code\Golang\Projects\cortex\temp\be3f0530-51e7-4911-a694-d04f4d180357
