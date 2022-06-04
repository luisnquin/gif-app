from pydantic import BaseModel
from typing import List, Tuple
from datetime import date


class Profile(BaseModel):
    job: str
    company: str
    ssn: str
    residence: str
    current_location: Tuple
    blood_group: str
    website: List[str]
    username: str
    name: str
    sex: str
    address: str
    mail: str
    birthdate: date
