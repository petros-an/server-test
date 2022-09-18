package main

import (
	"math/rand"
)

type PlayerId string

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

type Character struct {
	RigidBody     RigidBody2D
	Tag           string
	MoveDirection Vector2D
	speed         float64
	Color         RGBColor
}

func newCharacter(position Vector2D, color RGBColor) *Character {
	c := Character{}
	c.RigidBody.LocalPosition = position
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

func RandomColor() RGBColor {
	return RGBColor{
		R: uint8(rand.Intn(155) + 100),
		B: uint8(rand.Intn(155) + 100),
		G: uint8(rand.Intn(155) + 100),
	}
}
