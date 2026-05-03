from rq import Queue
from src.worker.connection import redis_conn

q = Queue("cortex", default_timeout=3600, connection=redis_conn)


print(q.connection, q.name)
