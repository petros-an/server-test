package main

import (
	"math/rand"
	"time"

	"github.com/petros-an/server-test/api"
	"github.com/petros-an/server-test/game"
)

func main() {

	g := game.New()

	rand.Seed(time.Now().UnixNano())
	go g.Run()
	api.Run(g)

}
