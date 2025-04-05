from typing import List

from pydantic import BaseModel


class Source(BaseModel):
    url: str


class SearchData(BaseModel):
    headline: str
    content: str | None
    url: str | None


class SearchResponse(BaseModel):
    result: str
    score: int
    sources: List[Source]
