from qdrant_client import QdrantClient
from worker.connection import redis_conn

# print()

# Connect to Qdrant
client = QdrantClient("http://localhost:6333")

# Get list of all collections
collections = client.get_collections().collections
collection_names = [c.name for c in collections]

# Delete each collection
for name in collection_names:
    client.delete_collection(collection_name=name)
    print(f"Deleted: {name}")


redis_conn.flushdb()
