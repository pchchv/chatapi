import datetime
from pydantic import BaseModel


class User(BaseModel):
    id: str
    username: str
    created_at: datetime


class Chat(BaseModel):
    id: str
    name: str
    users: [User.id]
    created_at: datetime


class Message(BaseModel):
    id: str
    chat: Chat.id
    author: User.id
    text: str
    created_at: datetime
