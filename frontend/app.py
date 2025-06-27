from fastapi import FastAPI
from fastapi.responses import FileResponse
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()


app.add_middleware(
  CORSMiddleware,
  allow_origins = ["127.0.0.1", "localhost"],
  allow_methods = ["*"],
  allow_headers = ["*"],
)

@app.get("/")
def index():
    return FileResponse("templates/index.html")


@app.get("/login")
def login():
    return FileResponse("templates/login.html")

app.mount("/static", StaticFiles(directory="static"), name="static")

# pip install fastapi uvicorn в виртуальной среде
# запуск сервера: uvicorn app:app --port 8001

