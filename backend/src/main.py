from fastapi import FastAPI
from src.api import health, ingest, job_result

app = FastAPI()


app.include_router(health.router)
app.include_router(ingest.router)
app.include_router(job_result.router)
