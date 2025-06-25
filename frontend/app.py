from fastapi import FastAPI
from fastapi.responses import FileResponse
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

@app.get("/")
def index():
    return FileResponse("templates/index.html")

app.mount("/static", StaticFiles(directory="static"), name="static")

# pip install fastapi uvicorn в виртуальной среде
# запуск сервера: uvicorn app:app --port 8001

