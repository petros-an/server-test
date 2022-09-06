// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func checkOrigin(r *http.Request) bool {
	return true
}

type CharacterId string

type Character struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	vx int
	vy int
	Id CharacterId
}

func (c *Character) move() {
	c.X += c.vx
	c.Y += c.vy
}

type GameState struct {
	Characters []*Character
}

type StateInput struct {
	NewCharacter   *Character
	VelocityUpdate *VelocityUpdate
}

type VelocityUpdate struct {
	Vx int
	Vy int
}

func newCharacter(id CharacterId) *Character {
	character := Character{X: rand.Intn(300), Y: rand.Intn(300), Id: id}
	return &character
}

type Endpoint func(http.ResponseWriter, *http.Request)

func getEndpoint(
	stateReads chan GameState,
	stateInputs chan StateInput,
) Endpoint {
	endpoint := func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
			CheckOrigin:     checkOrigin,
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

func main() {

	stateReadChan := make(chan GameState, 100)
	stateUpdateChan := make(chan StateInput, 100)
	go gameStateMaintainer(stateReadChan, stateUpdateChan, nil)
	http.HandleFunc(
		"/state",
		getEndpoint(stateReadChan, stateUpdateChan),
	)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func gameStateMaintainer(
	output chan GameState,
	input chan StateInput,
	stopper chan bool,
) {
	gameState := GameState{
		Characters: []*Character{},
	}

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			//log.Println("[M] Sending new state")
			//log.Println(gameState.Characters)
			output <- gameState
		case stateUpdate := <-input:
			//log.Println("[M] Received state update:")
			gameState = applyStateUpdate(gameState, stateUpdate)
			//log.Println(gameState)
		default:
			//log.Println("[M] Sleeping")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func applyStateUpdate(oldState GameState, update StateInput) GameState {

	for _, char := range oldState.Characters {
		if char.Id == update.NewCharacter.Id {
			return oldState
		}
	}

	oldState.Characters = append(oldState.Characters, update.NewCharacter)
	return oldState
}
