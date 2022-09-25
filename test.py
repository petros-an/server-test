import websocket
import uuid
import json
import time

def listen():
#    url = 'ws://still-citadel-50381.herokuapp.com/state'
    url = 'ws://localhost:8080/state'
    print('connecting...')
    ws = websocket.create_connection(url)

    print('sending id')
    
    t1 = time.time()
    ws.send(str(uuid.uuid4()))

    ws.recv()
    t2 = time.time()
    
    print(t2-t1)

    return 
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
