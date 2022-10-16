package projectile

import (
	"math/rand"

	"github.com/petros-an/server-test/common/collider"
	"github.com/petros-an/server-test/common/collider/shape"
	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/rigidbody"
	transform "github.com/petros-an/server-test/common/tansform"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
	"github.com/petros-an/server-test/game/gameObject"
)

const DefaultProjectileSpeed = 50
const DefaultProjectileDamage = 13

type Projectile struct {
	gameObject.GameObjectBasic

	RigidBody rigidbody.RigidBody2D
	Color     color.RGBColor
	FiredBy   *character.Character
	Damage    float64
	Id        int
	Collider  collider.Collider2D
}

func New(
	firedBy *character.Character,
	position vector.Vector2D,
	direction vector.Vector2D,
) *Projectile {
	velocity := direction.Mul(
		DefaultProjectileSpeed,
	).Add(firedBy.MoveVelocity())
	velocity.SetMagnitude(DefaultProjectileSpeed)
	p := Projectile{
		Color:   firedBy.Color,
		Id:      rand.Intn(100000),
		FiredBy: firedBy,
		Damage:  DefaultProjectileDamage,
		RigidBody: rigidbody.New(
			position, vector.New(0.5, 0.5), direction, velocity,
		),
	}

	coll := collider.NewBasic(&p, &shape.Ellipse{})
	coll.OnCollideField = p.onCollide
	p.Collider = coll

	return &p
}

func (p *Projectile) GetType() gameObject.GameObjectType {
	return gameObject.Projectile
}

func (p *Projectile) GetTransform() transform.Transform2D {
	return p.RigidBody.Transform2D
}

func (p *Projectile) Update(dt float64) {
	p.RigidBody.Update(dt)
}

func (p *Projectile) CollidesWithCharacter(c *character.Character) bool {
	return c.Position().Sub(p.RigidBody.Position).MagnitudeSq() < utils.Pow2(c.RigidBody.Scale.X/2+p.RigidBody.Scale.Y/2)
}

func (p *Projectile) onCollide(gobj gameObject.GameObject) {
	if gobj.GetType() == gameObject.Character {
		c := gobj.(*character.Character)
		if p.FiredBy == c {
			return
		}
		if c.GetDamaged(p.Damage) {
			p.FiredBy.AddKill()
		}
		p.Destroy()
	}
}
