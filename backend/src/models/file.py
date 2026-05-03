from pydantic import BaseModel


class FileModel(BaseModel):
    file_no: int
    file_path: str
