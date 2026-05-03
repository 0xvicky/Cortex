from langchain_qdrant import QdrantVectorStore
from langchain_community.embeddings import HuggingFaceEmbeddings


def query_svc(user_query: str):
    print(user_query)

    embeddings = HuggingFaceEmbeddings(model_name="BAAI/bge-small-en")

    vector_store = QdrantVectorStore.from_existing_collection(
        embedding=embeddings,
        collection_name="cortex_vecdb",
        url="http://localhost:6333",
    )

    search_result = vector_store.similarity_search(query=user_query, k=10)
    print(search_result)
    return "ok"
