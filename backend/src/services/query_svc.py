from langchain_qdrant import QdrantVectorStore
from langchain_community.embeddings import HuggingFaceEmbeddings
from dotenv import load_dotenv
from src.services.llm import llm_query
import os

load_dotenv()


def build_context(docs):
    context = {"repo_id": None, "files": {}}

    for d in docs:
        meta = d.metadata
        file_path = meta["file_path"]
        repo_id = meta["repo_id"]
        start = meta.get("start_line", 0)
        end = meta.get("end_line", 0)

        if context["repo_id"] is None:
            context["repo_id"] = repo_id

        # init file
        if file_path not in context["files"]:
            context["files"][file_path] = {"file_path": file_path, "chunks": []}

        # dedup check
        existing = context["files"][file_path]["chunks"]
        if any(c["start"] == start for c in existing):
            continue

        context["files"][file_path]["chunks"].append(
            {"start": start, "end": end, "content": d.page_content}
        )

    # sort chunks inside each file
    for f in context["files"].values():
        f["chunks"].sort(key=lambda x: x["start"])

    # convert dict → list
    context["files"] = list(context["files"].values())

    return context


def context_to_prompt(context):
    prompt = ""

    for f in context["files"]:
        prompt += f"\n### FILE: {f['file_path']}\n"

        for c in f["chunks"]:
            prompt += f"\nLines {c['start']}-{c['end']}:\n"
            prompt += c["content"] + "\n"

    return prompt


def query_svc(user_query: str, job_id: str):
    print(user_query)

    embeddings = HuggingFaceEmbeddings(model_name="BAAI/bge-small-en")

    vector_store = QdrantVectorStore.from_existing_collection(
        embedding=embeddings,
        collection_name=job_id,
        url=os.getenv("QDRANT_URL") or "http://host.docker.internal:6333",
        api_key=os.getenv("QDRANT_API_KEY") or None,
    )

    search_result = vector_store.similarity_search(query=user_query, k=5)
    context_obj = build_context(search_result)
    context_prompt = context_to_prompt(context_obj)
    print(search_result)
    SYSTEM_PROMPT = f"""
        You are a senior software engineer deeply familiar with this codebase.

        Answer the user's question using the provided context.
        If the exact answer isn't available, reason from what is available and give your best engineering judgment — stay grounded in the codebase, never go off-topic.
        Keep responses concise and to the point.

        Context:
        {context_prompt}

        Question:
        {user_query}

        Answer:
        """

    output = llm_query(
        system_prompt=SYSTEM_PROMPT,
        user_query=user_query,
    )

    return output
