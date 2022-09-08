package main

import (
	"math/rand"
	"time"
)

type CharacterId string

type Character struct {
	Position Vector2D
	Velocity Vector2D
	Id       CharacterId
}

const VelMagnitude float64 = 10

func (c *Character) move() {
	c.Position.AddSelf(c.Velocity.Mul(DT * VelMagnitude))
}

func spawnNewCharacter(id CharacterId) *Character {
	rand.Seed(time.Now().UTC().UnixNano())
	character := Character{Position: Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}, Id: id}
	return &character
}
