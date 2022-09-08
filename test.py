import websocket
import uuid
import json

def listen():
    url = 'wss://still-citadel-50381.herokuapp.com:8080'
    # url = 'ws://localhost:8080/state'
    print('connecting...')
    ws = websocket.create_connection(url)

    print('sending id')
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
