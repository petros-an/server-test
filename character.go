package main

import (
	"math/rand"
	"time"
)

type CharacterId string

type RGBColor struct {
	R int8
	G int8
	B int8
}

type Character struct {
	Position Vector2D
	Velocity Vector2D
	Id       CharacterId
	Color    RGBColor
}

const VelMagnitude float64 = 10

func (c *Character) move() {
	c.Position.AddSelf(c.Velocity.Mul(DT * VelMagnitude))
}

func spawnNewCharacter(id CharacterId) *Character {
	rand.Seed(time.Now().UTC().UnixNano())
	character := Character{
		Position: Vector2D{
			X: rand.Float64()*80 - 40,
			Y: rand.Float64()*80 - 40,
		},
		Id:    id,
		Color: RandomColor(),
	}
	return &character
}

func RandomColor() RGBColor {
	return RGBColor{
		R: int8(rand.Intn(256)),
		B: int8(rand.Intn(256)),
		G: int8(rand.Intn(256)),
	}
}
