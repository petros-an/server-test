package main

import (
	"math/rand"
	"time"
)

type Player struct {
	Character *Character
	LastVital time.Time
}

func newPlayer(position Vector2D, color RGBColor) *Player {
	new := Player{
		Character: newCharacter(position, color),
		LastVital: time.Now(),
	}
	return &new
}

func spawnNewPlayer() *Player {
	rand.Seed(time.Now().UTC().UnixNano())
	newPlayer := newPlayer(Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}, RandomColor())
	return newPlayer
}
