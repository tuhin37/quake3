from fastapi import APIRouter, FastAPI
from fastapi.responses import JSONResponse
from app.constants import Constants
from app.model import Expected
from dotenv import load_dotenv
load_dotenv()

default_router = APIRouter()


@default_router.get("/")
async def sanity():
    return JSONResponse(content=Constants.DEFAULT_SERVER_STATUS)


@default_router.post("/set_game")
async def set_game(settings: Expected):
    return JSONResponse(settings.dict())


app = FastAPI()
app.include_router(default_router)
