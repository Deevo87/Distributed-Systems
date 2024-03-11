from fastapi import FastAPI, HTTPException
from enum import Enum
from pydantic import BaseModel
from typing import Dict
from collections import defaultdict

app=FastAPI()

@app.get("/")
async def root() :
    return {"message" : "Hello World"}

class Poll(BaseModel):
    id: int
    name: str
    question: str
    options: Dict[str, int] = defaultdict(int)

polls: Dict[int, Poll] = {}

@app.post("/poll")
async def create_poll(poll: Poll):
    if poll.id not in polls:
        polls[poll.id] = poll
        return True # poll created
    else:
        raise HTTPException(status_code=404, detail="There is poll with that id: " + poll.id)

@app.get("poll/{poll_id}/vote/{option}")
async def vote(poll_id: int, option: str):
    if poll_id not in polls:
        raise HTTPException(status_code=404, detail="Poll not found.")
    else:
        polls[poll_id].options[option] += 1
        return True # voted

@app.get("/poll")
async def show_polls():
    return polls

@app.get("/poll/{poll_id}")
async def show_poll(poll_id: int):
    if poll_id not in polls:
        raise HTTPException(status_code=404, detail="Poll not found.")
    else:
        return polls[poll_id]

@app.post("/poll/{poll_id}/vote/{option_name}")
async def add_option(poll_id: int, option_name: str):
    if poll_id not in polls:
        raise HTTPException(status_code=404, detail="Poll not found.")
    else:
        polls[poll_id].options[option_name] = 0
        return True #option added

@app.delete("/poll/delete/{poll_id}")
async def del_poll(poll_id: int):
    if poll_id not in polls:
        raise HTTPException(status_code=404, detail="Poll not found.")
    else:
        del polls[poll_id]
        return True

