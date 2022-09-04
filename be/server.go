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

var upgrader = websocket.Upgrader{} // use default options

func checkOrigin(r *http.Request) bool {
	return true
}

type Character struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameState struct {
	Characters []*Character
}

var gameState GameState
var playerMap map[int]*Character
var ID int = 0

func addCharToRandomPosition() *Character {
	character := Character{X: rand.Intn(300), Y: rand.Intn(300)}
	gameState.Characters = append(gameState.Characters, &character)
	return &character
}

func handleStateEndpointConnection(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin:     checkOrigin,
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	newCharacter := addCharToRandomPosition()
	playerMap[ID] = newCharacter
	serializedID := serializeJson(&ID)
	log.Println(ID)
	ID++
	err = c.WriteMessage(websocket.TextMessage, serializedID)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {

		message := serializeJson(&gameState)
		err = c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
		time.Sleep(1 * time.Second)

	}
}

func serializeGameState(gameState *GameState) []byte {
	type GameStateData struct {
		CharactersData []Character
	}

	gameStateData := GameStateData{
		CharactersData: []Character{},
	}
	return serializeJson(&gameStateData)
}

func serializeJson(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	gameState = GameState{
		Characters: []*Character{},
	}
	playerMap = make(map[int]*Character)
	http.HandleFunc("/state", handleStateEndpointConnection)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
