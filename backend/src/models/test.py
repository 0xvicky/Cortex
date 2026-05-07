from pydantic import BaseModel, Field


class TestRepo(BaseModel):
    repo_url: str = Field(alias="repoUrl")
    date: str = Field(alias="date")

    class Config:
        populate_by_name = True
        extra = "forbid"
