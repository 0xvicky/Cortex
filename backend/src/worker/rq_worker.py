from rq import SimpleWorker, Queue
from src.worker.connection import redis_conn

q = Queue("cortex", connection=redis_conn, default_timeout=-1)
w = SimpleWorker([q], connection=redis_conn)
w.work(burst=False)
