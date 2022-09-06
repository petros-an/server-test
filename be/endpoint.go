package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Endpoint func(http.ResponseWriter, *http.Request)

func getEndpoint(
	stateReads chan GameState,
	stateInputs chan StateInput,
) Endpoint {
	endpoint := func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			//log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		//.Println("[E] Handling connection ...")

		//log.Println("[E] Reading id from client")
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		var id CharacterId = CharacterId(message)

		log.Println("[E] Adding new character")
		log.Println(id)

		input := StateInput{
			NewCharacter: newCharacter(id),
		}

		//log.Println(stateUpdate)

		//log.Println("[E] Sending client state update")

		stateInputs <- input

		for newState := range stateReads {
			//log.Println("[E] Sending state to client")
			message := serializeJson(&newState)
			err = c.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}

	return endpoint
}

func serializeJson(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}
