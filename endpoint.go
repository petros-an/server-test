package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Endpoint func(http.ResponseWriter, *http.Request)

type DirectionInput struct {
	X float64
	Y float64
}

func readStateInputsFromConnection(inputsChan chan StateInput, conn *websocket.Conn, charId CharacterId) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		var input map[string]interface{}
		err = parseJson(message, &input)
		if err != nil {
			log.Println(err)
		} else {

			if _, ok := input["ping"]; ok {
				// sendPing(conn)
			}
			handleMessageRecieved(input, inputsChan, charId)
		}
	}
}

func sendPing(conn *websocket.Conn) {
	message := serializeJson([]byte(`{"ping":{}}`))
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("ping error:", err)
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
		// log.Println(string(message))
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

func handleMessageRecieved(parsed map[string]interface{}, inputsChan chan StateInput, charId CharacterId) {
	for k, v := range parsed {
		switch k {
		case "direction":
			var directionInput DirectionInput
			temp, _ := json.Marshal(v)
			json.Unmarshal(temp, &directionInput)
			inputsChan <- StateInput{
				VelocityUpdate: &PlayerMoveDirectionUpdate{
					CharacterId:   charId,
					MoveDirection: Vector2D{X: directionInput.X, Y: directionInput.Y},
				},
			}
			break
		}
	}

}
