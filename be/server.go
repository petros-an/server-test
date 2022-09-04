// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

import (
	"flag"
	"log"
	"net/http"
    "math/rand"
	"github.com/gorilla/websocket"
    "encoding/json"
    "time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func checkOrigin(r *http.Request) bool {
    return true
}

type CharacterData struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type State struct {
    CharacterData []CharacterData;
}

// {
//     "characterData" : [
//         {
//             "x": 1,
//             "y": 2,
//         }
//     ]
// }

var gameState State;

func addCharToRandomPosition(){
    d := CharacterData{X: rand.Intn(300), Y:rand.Intn(300)}
    gameState.CharacterData = append(gameState.CharacterData, d)
}

func handleStateEndpointConnection(w http.ResponseWriter, r *http.Request) {
    var upgrader = websocket.Upgrader{
        ReadBufferSize:  4096,
        WriteBufferSize: 4096,
        CheckOrigin: checkOrigin,
    }

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

    addCharToRandomPosition()

	for {

        message := serializeJson(&gameState)
        log.Println(string(message))
		err = c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
        time.Sleep(1*time.Second)
        
	}
}

func serializeJson(data interface{}) []byte {
    res, _ := json.Marshal(data)
    return res
}

func main() {
	flag.Parse()
	log.SetFlags(0)
    gameState = State {
        CharacterData: []CharacterData{
            CharacterData{
                X: 12,
                Y: 23,
            },

        },
    }
	http.HandleFunc("/state", handleStateEndpointConnection)
	log.Fatal(http.ListenAndServe(*addr, nil))
}