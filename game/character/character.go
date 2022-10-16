package character

import (
	"github.com/petros-an/server-test/common/collider"
	"github.com/petros-an/server-test/common/collider/shape"
	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	transform "github.com/petros-an/server-test/common/tansform"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/gameObject"
)

type Character struct {
	gameObject.GameObjectBasic

	RigidBody     rigidbody.RigidBody2D
	Tag           string
	MoveDirection vector.Vector2D
	Color         color.RGBColor
	Health        float64
	KillCount     uint
	Collider      collider.Collider2D
}

func (c *Character) GetType() gameObject.GameObjectType {
	return gameObject.Character
}

func (c *Character) GetTransform() transform.Transform2D {
	return c.RigidBody.Transform2D
}

func (c *Character) AddKill() {
	c.KillCount++
}

func RandomNew() *Character {
	c := Character{
		RigidBody: rigidbody.NewInRandomPosition(
			vector.Vector2D{X: 3, Y: 3},
			vector.Vector2D{X: 1, Y: 0},
			vector.Zero(),
		),
		Color:  color.Random(),
		Health: DefaultHealth,
	}

	c.Collider = collider.NewBasic(&c, &shape.Ellipse{})

	return &c
}

const DefaultSpeed float64 = 20
const DefaultHealth float64 = 100

func (c Character) Position() vector.Vector2D {
	return c.RigidBody.FinalPosition()
}

func (c Character) MoveVelocity() vector.Vector2D {
	return c.MoveDirection.Mul(DefaultSpeed)
}

func (c *Character) SetPosition(position vector.Vector2D) {
	c.RigidBody.SetPosition(position)
}

func (c *Character) Update(dt float64) {
	c.RigidBody.Update(dt)
}

func (c *Character) SetMoveDirection(direction vector.Vector2D) {
	c.MoveDirection = direction
	v := c.MoveDirection.Mul(DefaultSpeed)
	c.RigidBody.Velocity = vector.Zero()
	c.RigidBody.Velocity.AddSelf(v)
}

func (c *Character) GetDamaged(damage float64) bool {
	// check if already dead
	if c.Health == 0 {
		return false
	}
	c.Health = utils.Max(c.Health-damage, 0)
	if c.Health == 0 {
		c.Die()
		return true
	}
	return false
}

func (c *Character) Die() {
	c.Respawn()
}

func (c *Character) Respawn() {
	c.SetPosition(vector.RandomNew())
	c.Health = DefaultHealth
}
