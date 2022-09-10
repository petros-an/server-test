package main

import (
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")
var addr = ":" + getPort(os.Getenv("PORT"))

func main() {

	stateReadChan := make(chan GameState, 100)
	stateUpdateChan := make(chan StateInput, 100)

	go gameStateMaintainer(stateReadChan, stateUpdateChan, nil)

	http.HandleFunc(
		"/state",
		getEndpoint(stateReadChan, stateUpdateChan),
	)
	log.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getPort(port string) string {
	if port == "" {
		return "8080"
	}
	return port
}
