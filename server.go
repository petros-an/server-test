package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

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
