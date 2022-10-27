from models import User
from fastapi import FastAPI
from dotenv import load_dotenv
from utils import create_user, create_chat


app = FastAPI()
load_dotenv('.env')


@app.get('/')
async def root():
    return {'message': 'Chat server. Version 0.0.1', 'status': 'ok'}


@app.post('/user/add')
async def read_user(username: str):
    user = create_user(username)
    return user


@app.post('/chat/add')
async def read_chat(name: str, users: [User.id]):
    chat = create_chat(name, users)
    return chat
