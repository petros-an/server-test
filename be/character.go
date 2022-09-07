package main

import "math/rand"

type CharacterId string

type Character struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	vx int
	vy int
	Id CharacterId
}

func (c *Character) move() {
	c.X += c.vx
	c.Y += c.vy
}

func newCharacter(id CharacterId) *Character {
	character := Character{X: rand.Intn(400) - 200, Y: rand.Intn(400) - 200, Id: id}
	return &character
}
