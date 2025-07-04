import os

from fastapi import FastAPI
from fastapi.responses import FileResponse
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()


origins = [
    "http://localhost:8001",
    "http://127.0.0.1:8001",
    "https://strrv.ru",
    "https://www.strrv.ru",
]

STATIC_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)), "static")

print(STATIC_DIR)

app.add_middleware(
  CORSMiddleware,
  allow_origins = origins,
  allow_methods = ["*"],
  allow_headers = ["*"],
  allow_credentials = True,
)

@app.get("/")
def index():
    return FileResponse("templates/index.html")


@app.get("/login")
def login():
    return FileResponse("templates/login.html")

app.mount("/static", StaticFiles(directory=STATIC_DIR), name="static")

# pip install fastapi uvicorn в виртуальной среде
# запуск сервера: uvicorn app:app --port 8001

