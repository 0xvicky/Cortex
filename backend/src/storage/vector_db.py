from qdrant_client import QdrantClient
from qdrant_client.models import VectorParams, Distance
from langchain_community.embeddings import HuggingFaceEmbeddings
from langchain_community.vectorstores import Qdrant
from langchain_qdrant import QdrantVectorStore
from langchain_core.documents import Document
from typing import List
from src.models.chunks import ChunkModel


def embed_chunks(job_id, chunks: List[ChunkModel]):

    # client = QdrantClient(host="localhost", port=6333)
    print("in embed chunks")
    docs = []

    for chunk in chunks:
        docs.append(
            Document(
                page_content=chunk.content,
                metadata={
                    "file_path": chunk.file_path,
                    "repo_id": job_id,
                    "start_line": chunk.start_line,
                    "end_line": chunk.end_line,
                },
            )
        )

    embeddings = HuggingFaceEmbeddings(model_name="BAAI/bge-small-en")
    # print(embeddings)

    try:
        vector_store = QdrantVectorStore.from_documents(
            docs, embeddings, collection_name=job_id, url="http://localhost:6333"
        )

        return True
    except Exception as e:
        return False
    # print(vector_store)
