from src.models.chunks import ChunkModel
from pydantic import BaseModel


class JobModel(BaseModel):
    job_id: str
    user_id: str
    status: str
    chunks: list[ChunkModel] | None = None
