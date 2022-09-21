package main

import (
	"log"
	"net/http"
	"os"
)

var addr = ":" + getEnv("PORT", "8080")

func main() {

	outputChannel := make(chan OutputMessage, 100)
	inputChannel := make(chan InputMessage, 100)

	go gameStateMaintainer(outputChannel, inputChannel, nil)

	http.HandleFunc(
		"/state",
		getEndpoint(outputChannel, inputChannel),
	)

	http.HandleFunc(
		"/ping",
		ping,
	)
	log.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getEnv(name string, fallback string) string {
	val := os.Getenv(name)
	if val == "" {
		return fallback
	}
	return val
}
