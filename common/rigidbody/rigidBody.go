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
