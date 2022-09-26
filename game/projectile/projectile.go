package projectile

import (
	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
)

const DefaultProjectileSpeed = 20

type Projectile struct {
	RigidBody     rigidbody.RigidBody2D
	MoveDirection vector.Vector2D
	Speed         float64
	Color         color.RGBColor
	FiredBy       *character.Character
}

func New(
	position vector.Vector2D,
	direction vector.Vector2D,
) *Projectile {
	p := Projectile{
		Color:         color.Random(),
		Speed:         DefaultProjectileSpeed,
		MoveDirection: direction,
	}
	p.RigidBody.LocalPosition = position
	return &p
}

func (p *Projectile) Update(dt float64) {
	v := p.MoveDirection.Mul(p.Speed)
	p.RigidBody.Velocity.AddSelf(v)
	p.RigidBody.Update(dt)
	p.RigidBody.Velocity.SubSelf(v)
}
