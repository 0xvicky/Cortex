from pydantic import BaseModel, Field


class UserQuery(BaseModel):
    user_query: str = Field(alias="userQuery")
    job_id: str = Field(alias="jobId")

    class Config:
        populate_by_name = True
        extra = "forbid"
