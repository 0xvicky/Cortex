# app/worker/connection.pyf
from redis import Redis

redis_conn = Redis(host="localhost", port=6379, db=0)
