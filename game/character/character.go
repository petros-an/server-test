package character

import (
	"math/rand"

	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	"github.com/petros-an/server-test/common/vector"
)

type Character struct {
	RigidBody     rigidbody.RigidBody2D
	Tag           string
	MoveDirection vector.Vector2D
	speed         float64
	Color         color.RGBColor
}

func RandomNew() *Character {
	c := Character{}
	c.RigidBody.LocalPosition = vector.Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}
	c.Color = color.Random()
	c.speed = DefaultVelMagnitude
	return &c
}

const DefaultVelMagnitude float64 = 10

func (c Character) Position() vector.Vector2D {
	return c.RigidBody.Position()
}

func (c *Character) SetPosition(position vector.Vector2D) {
	c.RigidBody.SetPosition(position)
}

func (c *Character) Update(dt float64) {
	v := c.MoveDirection.Mul(c.speed)
	c.RigidBody.Velocity.AddSelf(v)
	c.RigidBody.Update(dt)
	c.RigidBody.Velocity.SubSelf(v)
}
