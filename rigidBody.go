package main

type RigidBody2D struct {
	Transform2D
	Velocity Vector2D
}

func (this *RigidBody2D) Update() {
	this.updateVelocity()
}

func (this *RigidBody2D) updateVelocity() {
	this.LocalPosition.AddSelf(this.Velocity.Mul(DT))
}
