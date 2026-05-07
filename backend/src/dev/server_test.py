from qdrant_client import QdrantClient
from dotenv import load_dotenv
import os

# print()
load_dotenv()


def qdrant_cleanup():
    # Connect to Qdrant
    client = QdrantClient(
        "https://e91d65f3-16d2-4142-8029-a515d041e422.us-west-1-0.aws.cloud.qdrant.io",
        check_compatibility=False,
        api_key=os.getenv("QDRANT_API_KEY"),
    )

    # Get list of all collections
    collections = client.get_collections().collections
    collection_names = [c.name for c in collections]

    # Delete each collection
    for name in collection_names:
        client.delete_collection(collection_name=name)
        print(f"Deleted: {name}")


# redis_conn.flushdb()


def main():
    qdrant_cleanup()


main()
