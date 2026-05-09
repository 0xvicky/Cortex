from langchain_qdrant import QdrantVectorStore
from langchain_core.documents import Document
from typing import List
from src.models.chunks import ChunkModel
import os
from dotenv import load_dotenv
from langchain_huggingface import HuggingFaceEndpointEmbeddings

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

    embeddings = HuggingFaceEndpointEmbeddings(
        model="BAAI/bge-small-en", huggingfacehub_api_token=os.getenv("HF_TOKEN")
    )

    try:
        vector_store = QdrantVectorStore.from_documents(
            docs,
            embeddings,
            collection_name=job_id,
            url=os.getenv("QDRANT_URL") or "http://host.docker.internal:6333",
            api_key=os.getenv("QDRANT_API_KEY") or None,
        )

        return "COMPLETED"
    except Exception as e:
        print(f"Error while embedding :{e}")
        return "PARTIAL_SUCCESS"
