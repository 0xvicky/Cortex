from redis import Redis
import os
from dotenv import load_dotenv

load_dotenv()

r = Redis.from_url(os.getenv("REDIS_URL"))

# Test 1: basic ping
print(r.ping())  # should print True

# Test 2: set and get
r.set("test_key", "hello")
print(r.get("test_key"))  # should print b'hello'

# Test 3: delete
r.delete("test_key")
print("All good!")
