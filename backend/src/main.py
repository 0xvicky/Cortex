from fastapi import FastAPI
from src.api import health, ingest

app = FastAPI()


app.include_router(health.router)
app.include_router(ingest.router)
