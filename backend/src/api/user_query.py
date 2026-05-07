from fastapi import APIRouter
from src.models.user_query import UserQuery
from src.services.query_svc import query_svc

router = APIRouter(prefix="/user-query", tags=["result"])


@router.post("/")
def query(r: UserQuery):
    res = query_svc(user_query=r.user_query, job_id=r.job_id)
    return {"res": res}
