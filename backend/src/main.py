from fastapi import FastAPI, Depends
from src.api import health, ingest, job_result, user_query, get_jobs, auth
from fastapi.middleware.cors import CORSMiddleware
from src.middleware.middleware import verify_jwt

app = FastAPI()


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Vite's default port
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Custom middleware
app.include_router(auth.router)
app.include_router(health.router)
app.include_router(ingest.router, dependencies=[Depends(verify_jwt)])
app.include_router(job_result.router, dependencies=[Depends(verify_jwt)])
app.include_router(user_query.router, dependencies=[Depends(verify_jwt)])
app.include_router(get_jobs.router, dependencies=[Depends(verify_jwt)])
