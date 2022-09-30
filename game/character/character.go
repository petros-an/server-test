package character

import (
	"math/rand"

	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/world"
)

type Character struct {
	RigidBody     rigidbody.RigidBody2D
	Tag           string
	MoveDirection vector.Vector2D
	speed         float64
	Color         color.RGBColor
	Health        float64
}

func RandomNew() *Character {
	c := Character{}
	c.RigidBody.LocalPosition = vector.Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}
	c.RigidBody.LocalScale = vector.Vector2D{X: 3, Y: 3}
	c.Color = color.Random()
	c.speed = DefaultSpeed
	c.Health = DefaultHealth
	return &c
}

const DefaultSpeed float64 = 20
const DefaultHealth float64 = 100

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

	c.SetPosition(
		world.RestrictPositionWithinBorder(c.Position()),
	)
}

func (c *Character) SetMoveDirection(direction vector.Vector2D) {
	c.MoveDirection = direction
}

func (c *Character) GetDamaged(damage float64) {
	if c.Health < 0 {
		return
	}
	c.Health -= damage
	if c.Health < 0 {
		c.Die()
		c.Health = 0
	}
}

func (c *Character) Die() {
}
