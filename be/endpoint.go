package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Endpoint func(http.ResponseWriter, *http.Request)

type ClientInput struct {
	Velocity struct {
		VX float64
		VY float64
	}
}

func readStateInputsFromConnection(inputsChan chan StateInput, conn *websocket.Conn, charId CharacterId) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		var input ClientInput
		err = parseJson(message, &input)
		if err != nil {
			log.Println(err)
		} else {
			inputsChan <- StateInput{
				VelocityUpdate: &VelocityUpdate{
					CharacterId: charId,
					Velocity:    Vector2D{X: input.Velocity.VX, Y: input.Velocity.VY},
				},
			}
		}
	}
}

func sendStateToConnection(stateReads chan GameState, conn *websocket.Conn) {
	for newState := range stateReads {
		//log.Println("[E] Sending state to client")
		message := serializeJson(&newState)
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func readIDFromConnection(conn *websocket.Conn) CharacterId {

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return ""
	}
	var id CharacterId = CharacterId(message)

	log.Println("[E] Adding new character")
	log.Println(id)

	return id
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getEndpoint(
	stateReads chan GameState,
	stateInputs chan StateInput,
) Endpoint {
	endpoint := func(w http.ResponseWriter, r *http.Request) {

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		characterId := readIDFromConnection(c)
		input := StateInput{
			NewCharacter: spawnNewCharacter(characterId),
		}

		stateInputs <- input

		go readStateInputsFromConnection(stateInputs, c, characterId)

		sendStateToConnection(stateReads, c)

	}

	return endpoint
}

func serializeJson(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}

func parseJson(data []byte, dst interface{}) error {
	err := json.Unmarshal(data, dst)
	return err
}
