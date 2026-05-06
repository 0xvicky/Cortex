from fastapi import APIRouter
from src.models.user_query import UserQuery

from src.services.query_svc import query_svc

from src.worker.rq_worker import q

router = APIRouter(prefix="/user-query", tags=["result"])


@router.post("/")
def query(r: UserQuery):
    # print(job_id)
    # print("in res")
    res = query_svc(user_query=r.user_query, job_id=r.job_id)
    return {"res": res}
