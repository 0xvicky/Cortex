from rq import Queue
from src.worker.connection import redis_conn

q = Queue("cortex", connection=redis_conn)


print(q.connection, q.name)
