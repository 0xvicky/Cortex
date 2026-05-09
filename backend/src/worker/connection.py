# app/worker/connection.pyf
from redis import Redis
import os
from dotenv import load_dotenv

load_dotenv()
redis_conn = Redis.from_url(os.getenv("REDIS_URL"))
# redis_conn = Redis(host="host.docker.internal", port=6379)

# redis_conn = Redis(
#     host="redis-15347.crce206.ap-south-1-1.ec2.cloud.redislabs.com",
#     port=15347,
#     decode_responses=True,
#     username="default",
#     password="*******",
# )
