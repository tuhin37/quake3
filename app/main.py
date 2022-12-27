from fastapi import APIRouter, FastAPI
from fastapi.responses import JSONResponse
from app.constants import Constants
from app.model import Expected
from dotenv import load_dotenv
import os
from app.generators.file_generator import create_file
import daemon
from app.executors.game_executor import GameExecutor
import os

load_dotenv()

default_router = APIRouter()
RAM = os.getenv("RAM") if os.getenv("RAM") else 128
PORT = os.getenv("PORT") if os.getenv("PORT") else 27960


def create_autoexec():
    """this will create the autoexec file"""
    path = "/usr/local/games/quake3/baseq3/autoexec.cfg"
    if not os.path.exists(path):
        # we are not inside docker, most probably testing.
        path = "./target"
        os.makedirs(path, exist_ok=True)
        path = os.path.join(path, "autoexec.cfg")
    target_str = f"""    set vm_game 2           // I have no idea what this shit is
                        set vm_cgame 2          // Nope
                        set vm_ui 2             // Nada
                        set dedicated 1         // Dedicated server but not announced
                        set com_hunkmegs {RAM}    // How much RAM for your server
                        set net_port {PORT}      // The network port
                """
    with open(path, "w") as writer:
        writer.write(target_str)


@default_router.get("/")
async def sanity():
    return JSONResponse(content=Constants.DEFAULT_SERVER_STATUS)


@default_router.post("/set_game")
async def set_game(settings: Expected):
    settings_file = create_file(settings)
    GameExecutor().run_game(settings_file)
    return JSONResponse(settings.dict())


create_autoexec()
app = FastAPI()
app.include_router(default_router)
