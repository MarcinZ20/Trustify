from fastapi import FastAPI

from .models.models import SearchData
from .processing.processing import process

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello from Trustify AI service!"}


@app.get("/health")
async def check_health():
    return {"message": "healthy"}


@app.post("/search")
async def search(payload: SearchData):
    logging.info(payload)
    result = process(payload)

    if result == None:
        return {"message": "ai-service-error"}

    return {
        "message": "Success",
        "result": result.result,
        "score": result.score,
        "sources": result.sources,
    }
