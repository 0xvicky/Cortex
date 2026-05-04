from langchain_qdrant import QdrantVectorStore
from langchain_community.embeddings import HuggingFaceEmbeddings
from langchain_groq import ChatGroq
from langchain.agents import create_agent
from langchain_core.messages import SystemMessage, HumanMessage
from dotenv import load_dotenv

load_dotenv()


def llm_query(system_prompt, user_query):
    llm = ChatGroq(
        model="qwen/qwen3-32b",
        temperature=0,
        max_tokens=None,
        reasoning_format="parsed",
        timeout=None,
        max_retries=2,
        # other params...
    )

    response = llm.invoke(
        [SystemMessage(content=system_prompt), HumanMessage(content=user_query)]
    )

    # agent = create_agent(
    #     llm,
    #     system_prompt=system_prompt,
    # )

    # result = agent.invoke({"messages": [{"role": "user", "content": user_query}]})

    # output = result["messages"][-1].content
    # return output
    return response.content
