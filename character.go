package main

import (
	"math/rand"
	"time"
)

type CharacterId string

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

type Character struct {
	RigidBody     RigidBody2D
	MoveDirection Vector2D
	speed         float64
	Id            CharacterId
	Color         RGBColor
}

func newCharacter(position Vector2D, id CharacterId, color RGBColor) *Character {
	c := Character{}
	c.RigidBody.LocalPosition = position
	c.Id = id
	c.Color = color
	c.speed = DefaultVelMagnitude
	return &c
}

const DefaultVelMagnitude float64 = 10

func (this Character) Position() Vector2D {
	return this.RigidBody.Position()
}

func (this *Character) SetPosition(position Vector2D) {
	this.RigidBody.SetPosition(position)
}

func (this *Character) Update() {
	v := this.MoveDirection.Mul(this.speed)
	this.RigidBody.Velocity.AddSelf(v)
	this.RigidBody.Update()
	this.RigidBody.Velocity.SubSelf(v)
}

func spawnNewCharacter(id CharacterId) *Character {
	rand.Seed(time.Now().UTC().UnixNano())
	character := newCharacter(Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}, id, RandomColor())
	return character
}

func RandomColor() RGBColor {
	return RGBColor{
		R: uint8(rand.Intn(155) + 100),
		B: uint8(rand.Intn(155) + 100),
		G: uint8(rand.Intn(155) + 100),
	}
}
