package main

import (
	"math/rand"
	"time"
)

type CharacterId string

type Character struct {
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	VX float32
	VY float32
	Id CharacterId
}

const VelMagnitude float32 = 10

func (c *Character) move() {
	c.X += c.VX * DT * VelMagnitude
	c.Y += c.VY * DT * VelMagnitude
}

func newCharacter(id CharacterId) *Character {
	rand.Seed(time.Now().UTC().UnixNano())
	character := Character{X: rand.Float32()*80 - 40, Y: rand.Float32()*80 - 40, Id: id}
	return &character
}
