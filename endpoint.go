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

func readStateInputsFromConnection(inputChannel chan InputMessage, conn *websocket.Conn, charId PlayerId) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		var parsedMessage map[string]interface{}
		err = parseJson(message, &parsedMessage)
		if err != nil {
			log.Println(err)
		} else {
			handleMessageRecieved(parsedMessage, inputChannel, charId)
		}
	}
}

func sendStateToConnection(stateReads chan OutputMessage, conn *websocket.Conn) {

	// loops every time stateReads is set
	for newState := range stateReads {
		switch newState.Type {
		case O_STATE:
			//log.Println("[E] Sending state to client")
			message := serializeJson(&(newState.CurrentState))
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("write:", err)
				return
			}
			// log.Println(string(message))

		default:
		}

	}
}

func readIDFromConnection(conn *websocket.Conn) PlayerId {

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return ""
	}
	var id PlayerId = PlayerId(message)

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
	outputChannel chan OutputMessage,
	inputChannel chan InputMessage,
) Endpoint {
	endpoint := func(w http.ResponseWriter, r *http.Request) {

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		newPlayerId := readIDFromConnection(c)

		Input := InputMessage{
			Type:      I_NEW,
			PlayerId:  newPlayerId,
			NewPlayer: spawnNewPlayer(),
		}

		Input.NewPlayer.Character.Tag = string(newPlayerId)

		inputChannel <- Input

		go readStateInputsFromConnection(inputChannel, c, newPlayerId)

		sendStateToConnection(outputChannel, c)

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

func handleMessageRecieved(parsedMessage map[string]interface{}, inputChannel chan InputMessage, charId PlayerId) {
	for k, v := range parsedMessage {
		switch k {
		case "direction":
			var directionInput DirectionInput
			temp, _ := json.Marshal(v)
			json.Unmarshal(temp, &directionInput)
			inputChannel <- InputMessage{
				Type:      I_DIRECTION,
				PlayerId:  charId,
				Direction: newVector2D(directionInput.X, directionInput.Y),
			}
		}
	}

}

func ping(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("(ping)upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		message := serializeJson(map[string]interface{}{})
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("error during ping sending:", err)
			return
		}
	}

}
