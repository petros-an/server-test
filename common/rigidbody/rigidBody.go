package rigidbody

import (
	transform "github.com/petros-an/server-test/common/tansform"
	"github.com/petros-an/server-test/common/vector"
)

type RigidBody2D struct {
	transform.Transform2D
	Velocity vector.Vector2D
}

func (rb *RigidBody2D) Update(dt float64) {
	rb.updateVelocity(dt)
}

func (rb *RigidBody2D) updateVelocity(dt float64) {
	rb.Position.AddSelf(rb.Velocity.Mul(dt))
}

func New(
	position vector.Vector2D,
	scale vector.Vector2D,
	rotation vector.Vector2D,
	velocity vector.Vector2D,
) RigidBody2D {
	return RigidBody2D{
		Transform2D: transform.New(
			position, scale, rotation,
		),
		Velocity: velocity,
	}
}

func NewInRandomPosition(
	scale vector.Vector2D,
	rotation vector.Vector2D,
	velocity vector.Vector2D,
) RigidBody2D {
	return New(
		vector.RandomNew(), scale, rotation, velocity,
	)
}
