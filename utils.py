import datetime
from models import User, Chat, Message


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
    return chat


def create_message(chat: Chat.id, author: User.id, text: str):
    message = Message
    # TODO: Generate id
    message.chat = chat
    message.author = author
    message.text = text
    message.created_at = datetime.datetime.now()
    #  TODO: Add message to db
    return message


def get_chats(user: User.id):
    chats = []
    # TODO: Get all chats in which the user is a member
    return chats


def get_messages(chat: Chat.id):
    messages = []
    # TODO: Get all messages from chat
    return messages
