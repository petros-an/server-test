import websocket
import uuid
import json

def listen():
    ws = websocket.create_connection("ws://localhost:8080/state")

    ws.send(str(uuid.uuid4()))

    while True:
        print(ws.recv())
        data = {
                "Velocity":  {
                    "VX": 1, "VY":1
                }
        }
        msg = json.dumps(data)
        ws.send(msg)
listen()
