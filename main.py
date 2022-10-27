from fastapi import FastAPI
from dotenv import load_dotenv


app = FastAPI()
load_dotenv('.env')


@app.get('/')
async def root():
    return {'message': 'Chat server. Version 0.0.1', 'status': 'ok'}
