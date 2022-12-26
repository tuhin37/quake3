
from fastapi import APIRouter, FastAPI
from fastapi.responses import JSONResponse
from app.constants import Constants

default_router = APIRouter()

@default_router.get("/")
async def sanity():
    return JSONResponse(content=Constants.DEFAULT_SERVER_STATUS)


app = FastAPI()
app.include_router(default_router)