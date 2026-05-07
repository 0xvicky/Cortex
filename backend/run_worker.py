from rq import SimpleWorker
from src.worker.rq_worker import q
from src.worker.connection import redis_conn

w = SimpleWorker([q], connection=redis_conn)
w.work(burst=False)
print("this")
