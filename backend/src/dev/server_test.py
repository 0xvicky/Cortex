from qdrant_client import QdrantClient
from dotenv import load_dotenv
import os
import redis

# print()
load_dotenv()


def qdrant_cleanup():
    # Connect to Qdrant
    client = QdrantClient(
        "https://e91d65f3-16d2-4142-8029-a515d041e422.us-west-1-0.aws.cloud.qdrant.io",
        check_compatibility=False,
        api_key=os.getenv("QDRANT_API_KEY"),
    )
    try:

        # client = QdrantClient(
        #     "http://localhost:6333",
        #     check_compatibility=False,
        # )

        # Get list of all collections
        collections = client.get_collections().collections
        collection_names = [c.name for c in collections]
        print("deleting...")
        # Delete each collection
        for name in collection_names:
            client.delete_collection(collection_name=name)
            print(f"Deleted: {name}")
    except Exception as e:
        print(f"Error connecting to Qdrant: {e}")


# redis_conn.flushdb()


def redis_cleanup():

    r = redis.Redis.from_url(os.getenv("REDIS_URL"))
    # r = redis.Redis(host="host.docker.internal", port=6379)
    print(r)
    r.flushdb()


def main():
    qdrant_cleanup()
    redis_cleanup()


main()
# Warning: You are sending unauthenticated requests to the HF Hub. Please set a HF_TOKEN to enable higher rate limits and faster downloads.
