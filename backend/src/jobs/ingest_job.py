# app/jobs/ingest_job.py

from src.services.ingest import ingestr_svc
from src.storage.storage import update_status, store_chunks
from src.storage.vector_db import embed_chunks


def run_pipeline(job_id: str, repo_url: str, user_id: str):
    try:
        print("in pipeline")

        update_status(job_id, "RUNNING")

        chunks = ingestr_svc(repo_url)  # ✅ no asyncio.run
        # store chunks in redisdb
        print("CHUNKS BEFORE STORE:", len(chunks))
        store_chunks(job_id, chunks)
        # store the chunks in vector db for rag retrieval
        embed_chunks(job_id, chunks)
        update_status(job_id, "COMPLETED")

    except Exception as e:
        update_status(job_id, "FAILED")
        raise e
