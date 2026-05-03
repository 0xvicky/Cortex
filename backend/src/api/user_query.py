from fastapi import APIRouter, UploadFile, File, Form, HTTPException
from src.models.user_query import UserQuery
from src.services.ingest import ingestr_svc
from src.services.query_svc import query_svc
from typing import Optional
from src.worker.queue import q
from src.storage.storage import get_job
import uuid

router = APIRouter(prefix="/user-query", tags=["result"])


@router.get("/")
def query(r: UserQuery):
    # print(job_id)
    # print("in res")
    res = query_svc(user_query=r.user_query)
    return {"res": res}
