import datetime
from models import User


def create_user(name: str):
    user = User
    # TODO: Generate id
    user.username = name
    user.created_at = datetime.datetime.now()
    # TODO: Add user to db
    return user
