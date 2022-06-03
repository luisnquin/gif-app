from pydantic import BaseModel, Field
from typing import List


class Internal(BaseModel):
    port: str
    api_key: List[str] = Field(..., alias='apiKey')
    token_expiration_time: int = Field(..., alias='tokenExpirationTime')
    email_regex: str = Field(..., alias='emailRegex')


class Database(BaseModel):
    time_out: int = Field(..., alias='timeOut')
    in_local_dsn: str = Field(..., alias='inLocalDsn')
    in_container_dsn: str = Field(..., alias='inContainerDsn')


class Configuration(BaseModel):
    internal: Internal
    database: Database
