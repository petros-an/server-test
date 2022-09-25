package main

import (
	"github.com/petros-an/server-test/api"
	"github.com/petros-an/server-test/game"
)

func main() {

	g := game.New()

	go g.Run()
	api.Run(g)

}
