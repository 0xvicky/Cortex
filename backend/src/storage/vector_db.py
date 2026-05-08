from langchain_community.embeddings import HuggingFaceEmbeddings
from langchain_qdrant import QdrantVectorStore
from langchain_core.documents import Document
from typing import List
from src.models.chunks import ChunkModel
import os
from dotenv import load_dotenv

load_dotenv()


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

    try:
        vector_store = QdrantVectorStore.from_documents(
            docs,
            embeddings,
            collection_name=job_id,
            url=os.getenv("QDRANT_URL") or "http://host.docker.internal:6333",
            api_key=os.getenv("QDRANT_API_KEY") or None,
        )

        return True
    except Exception as e:
        return False
