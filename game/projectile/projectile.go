package projectile

import (
	"math/rand"

	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
	gameobject "github.com/petros-an/server-test/game/gameObject"
	"github.com/petros-an/server-test/game/world"
)

const DefaultProjectileSpeed = 50
const DefaultProjectileDamage = 15

type Projectile struct {
	RigidBody rigidbody.RigidBody2D
	Color     color.RGBColor
	FiredBy   *character.Character
	Damage    float64
	Id        int
}

func New(
	firedBy *character.Character,
	position vector.Vector2D,
	direction vector.Vector2D,
) *Projectile {
	p := Projectile{
		Color:   firedBy.Color,
		Id:      rand.Intn(100000),
		FiredBy: firedBy,
		Damage:  DefaultProjectileDamage,
	}
	p.RigidBody.SetPosition(position)
	p.RigidBody.SetScale(vector.New(0.5, 0.5))
	p.RigidBody.SetRotation(direction)
	p.RigidBody.Velocity = direction.Mul(DefaultProjectileSpeed)
	return &p
}

func (p *Projectile) Update(dt float64) {
	p.RigidBody.Update(dt)
}

func (p *Projectile) IsOutsideWorld() bool {
	return world.IsOutsideWorld(p.RigidBody.Position())
}

func (p *Projectile) CollidesWith(o gameobject.GameObject) bool {
	return false
}

func (p *Projectile) CollidesWithCharacter(c *character.Character) bool {
	return c.Position().Sub(p.RigidBody.LocalPosition).MagnitudeSq() < utils.Pow2(c.RigidBody.LocalScale.X/2+p.RigidBody.LocalScale.Y/2)
}
