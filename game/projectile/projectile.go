package projectile

import (
	"math/rand"

	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
)

const DefaultProjectileSpeed = 50

type Projectile struct {
	RigidBody rigidbody.RigidBody2D
	Color     color.RGBColor
	FiredBy   *character.Character
	Id        int
}

func New(
	firedBy *character.Character,
	position vector.Vector2D,
	direction vector.Vector2D,
) *Projectile {
	p := Projectile{
		Color:   color.Random(),
		Id:      rand.Intn(100000),
		FiredBy: firedBy,
	}
	p.RigidBody.LocalPosition = position
	p.RigidBody.Velocity = direction.Mul(DefaultProjectileSpeed)
	p.RigidBody.LocalScale = vector.New(0.5, 0.5)
	return &p
}

func (p *Projectile) Update(dt float64) {
	p.RigidBody.Update(dt)
}
