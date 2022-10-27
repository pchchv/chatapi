from fastapi import FastAPI
from utils import create_user
from dotenv import load_dotenv

app = FastAPI()
load_dotenv('.env')


@app.get('/')
async def root():
    return {'message': 'Chat server. Version 0.0.1', 'status': 'ok'}


@app.post('/user/add')
async def read_user(username: str):
    user = create_user(username)
    return user
