from fastapi import FastAPI
from models import User, Chat
from dotenv import load_dotenv
from utils import create_user, create_chat, create_message, get_chats, get_messages


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


@app.post('/message/add')
async def read_message(chat: Chat.id, author: User.id, text: str):
    message = create_message(chat, author, text)
    return message


@app.get('/chat')
async def read_chats(user: User.id):
    chats = get_chats(user)
    return chats


@app.get('/message')
async def read_messages(chat: Chat.id):
    messages = get_messages(chat)
    return messages
