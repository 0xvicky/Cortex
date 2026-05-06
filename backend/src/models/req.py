from pydantic import BaseModel, Field, HttpUrl
from typing import Optional


class RequestRepo(BaseModel):
    repo_url: HttpUrl = Field(alias="repoUrl")
    user_id: str = Field(alias="userId")

    class Config:
        populate_by_name = True
        extra = "forbid"
