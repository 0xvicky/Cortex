from pydantic import BaseModel, Field, HttpUrl
from typing import Optional


class TestRepo(BaseModel):
    repo_url: str = Field(alias="repoUrl")
    date: str = Field(alias="date")

    class Config:
        populate_by_name = True
        extra = "forbid"
