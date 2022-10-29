<div align="center">

# chatAPI

# A chat server that provides an HTTP API to handle user chats and messages

</div>

## HTTP Methods

```
"GET" / — Checking the server connection

    example: 
        "GET" :9000/
```

```
"GET" /ping — Checking the server connection

    example: 
        "GET" :9000/ping
```

```
"POST" /user — Add a new user
    options:
        name — User name 

    example: 
        "POST" :9000/user/add?name=Jack
```

```
"POST" /chat — Create a new chat. Need JSON body

    example: 
        "POST" :9000/chat
```

```json
{
    "name" : "Alex's birthday",
    "users": ["3333x45if", "454x8548o", "112y5641654"]
}
```

```
"POST" /message — Send a message to the chat. Need JSON body

    example: 
        "POST" :9000/message
```

```json
{
    "chat" : "xx45499e0",
    "author": "321x4445d",
    "text": "Hi there"
}
```

```
"GET" /user/chat — Get a list of the user's chats
    options:
        user — User id

    example: 
        "GET" :9000/user/chat?user=333x541i
```

```
"GET" /chat/messages — Get a list of chat messages
    options:
        chat — Chat id

    example: 
        "GET" :9000/chat/messages?chat=xx000fd1
```

```
"DELETE" /chat — Delete chat by id
    options:
        chat — Chat id

    example: 
        "DELETE" :9000/chat?chat=xx000fd1
```

```
"DELETE" /message — Delete message by id
    options:
        chat — Chat id
        message — Message id

    example: 
        "DELETE" :9000/chat?chat=xx000fd1,message=00031
```

```
"DELETE" /chat/all — Delete all chats

    example: 
        "DELETE" :9000/
```
