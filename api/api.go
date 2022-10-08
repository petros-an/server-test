package api

import (
	"log"
	"net/http"

	"github.com/petros-an/server-test/api/gameinfo"
	"github.com/petros-an/server-test/api/ping"
	"github.com/petros-an/server-test/api/state"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/game"
)

func Run(
	g *game.Game,
) {

	addr := ":" + utils.GetEnv("PORT", "8080")

	http.HandleFunc(
		"/state",
		state.GetEndpoint(g),
	)

	http.HandleFunc(
		"/ping",
		ping.Endpoint,
	)

	http.HandleFunc(
		"/gameinfo",
		gameinfo.GetEndpoint(g),
	)

	log.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
