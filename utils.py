import datetime
from models import User,  Chat


def create_user(name: str):
    user = User
    # TODO: Generate id
    user.username = name
    user.created_at = datetime.datetime.now()
    # TODO: Add user to db
    return user


def create_chat(name: str, users: [User.id]):
    chat = Chat
    # TODO: Generate id
    chat.name = name
    chat.users = users
    chat.created_at = datetime.datetime.now()
    # TODO: Add chat to db
