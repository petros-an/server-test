package main

import (
	"math/rand"
)

type CharacterId string

type Character struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	VX int
	VY int
	Id CharacterId
}

func (c *Character) move() {
	c.X += c.VX
	c.Y += c.VY
}

func newCharacter(id CharacterId) *Character {
	character := Character{X: rand.Intn(400) - 200, Y: rand.Intn(400) - 200, Id: id}
	return &character
}
