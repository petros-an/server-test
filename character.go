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
	RigidBody RigidBody2D
	Id        CharacterId
	Color     RGBColor
}

func newCharacter(position Vector2D, id CharacterId, color RGBColor) *Character {
	c := Character{}
	c.RigidBody.Position = position
	c.Id = id
	c.Color = color
	return &c
}

const VelMagnitude float64 = 10

func (this *Character) Update() {
	this.RigidBody.Update()
}

func spawnNewCharacter(id CharacterId) *Character {
	rand.Seed(time.Now().UTC().UnixNano())
	character := newCharacter(Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}, id, RandomColor())
	return character
}

func RandomColor() RGBColor {
	return RGBColor{
		R: int8(rand.Intn(256)),
		B: int8(rand.Intn(256)),
		G: int8(rand.Intn(256)),
	}
}
