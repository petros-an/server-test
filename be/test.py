import websocket
import uuid
def listen():
    ws = websocket.create_connection("ws://localhost:8080/state")

    ws.send(str(uuid.uuid4()))

    while True:
        print(ws.recv())
 
listen()
