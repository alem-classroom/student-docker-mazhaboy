import os
from typing import Optional
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def handler_root():
    return {"Hello": os.getenv("APP_TYPE")}

@app.get("/items/{item_id}")
def handler_items(item_id: int, q: Optional[str] = None):
    return {"item_id": item_id, "q": q}
