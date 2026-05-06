from pydantic import BaseModel


class ChunkModel(BaseModel):
    chunk_no: int
    file_no: int
    file_path: str
    content: str
    start_line: int
    end_line: int
