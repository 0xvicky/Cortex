from fastapi import FastAPI
from src.api import health, ingest, job_result, user_query, get_jobs
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Vite's default port
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
app.include_router(health.router)
app.include_router(ingest.router)
app.include_router(job_result.router)
app.include_router(user_query.router)
app.include_router(get_jobs.router)
