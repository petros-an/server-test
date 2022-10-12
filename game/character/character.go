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
	"github.com/petros-an/server-test/game/world"
)

type Character struct {
	RigidBody     rigidbody.RigidBody2D
	toDestroy     bool
	Tag           string
	MoveDirection vector.Vector2D
	Color         color.RGBColor
	Health        float64
	KillCount     uint
	Collider      *collider.Collider2D
}

func (c *Character) GetType() gameObject.GameObjectType {
	return gameObject.Character
}

func (c *Character) GetTransform() transform.Transform2D {
	return c.RigidBody.Transform2D
}

func (c *Character) ToDestroy() bool {
	return c.toDestroy
}

func (c *Character) Destroy() {
	c.toDestroy = true
}

func (c *Character) AddKill() {
	c.KillCount++
}

func RandomNew() *Character {
	c := Character{}
	c.RigidBody.Position = vector.RandomNew()
	c.RigidBody.Scale = vector.Vector2D{X: 3, Y: 3}
	c.Color = color.Random()
	c.Health = DefaultHealth
	v := c.MoveDirection.Mul(DefaultSpeed)
	c.RigidBody.Velocity.AddSelf(v)

	c.Collider = collider.New(&c, &shape.Ellipse{})
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

	c.SetPosition(
		world.RestrictPositionWithinBorder(c.Position(), c.RigidBody.Scale.Div(2)),
	)
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
