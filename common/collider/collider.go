package collider

import (
	"github.com/petros-an/server-test/common/collider/shape"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/game/gameObject"
)

type Collider2D struct {
	GameObject gameObject.GameObject
	Shape      shape.Shape2D

	PositionFromTrasnform bool
	ScaleFromTrasnform    bool
	RotationFromTrasnform bool

	OnCollide func(other gameObject.GameObject)
}

func New(gameObject gameObject.GameObject, shape shape.Shape2D) *Collider2D {
	coll := Collider2D{
		GameObject: gameObject,
		Shape:      shape,

		PositionFromTrasnform: true,
		ScaleFromTrasnform:    true,
		RotationFromTrasnform: true,
	}
	return &coll
}

func (this Collider2D) PrepareCollider() {

	this.Shape.SetPSR2D(this.GameObject.GetTransform().PSR2D, this.PositionFromTrasnform, this.ScaleFromTrasnform, this.RotationFromTrasnform)
	if pol, isPol := this.Shape.(*shape.Polygon); isPol {
		pol.PointsAreTransformed = false
	}
}

func (this Collider2D) IsCollidedWithCollider(other *Collider2D) bool {
	return this.Shape.OverlapsWithShape(other.Shape)
}

func (this Collider2D) IsCollidedWithShape(shape shape.Shape2D) bool {
	return this.Shape.OverlapsWithShape(shape)
}

func (this Collider2D) CanColliderWith(other *Collider2D) bool {
	return true
}

func (this Collider2D) ShouldntExist() bool {
	return this.GameObject.ToDestroy()
}

func CleanColliderSlice(colls *[]*Collider2D) {
	len := len(*colls)
	for i := len - 1; i >= 0; i-- {
		if (*colls)[i].ShouldntExist() {
			utils.RemoveElementFromSliceAtIndex(colls, i)
		}
	}
}

func CheckCollidersCollisions(colls *[]*Collider2D) {

	CleanColliderSlice(colls)

	for _, coll := range *colls {
		coll.PrepareCollider()
	}

	for _, coll1 := range *colls {
		if coll1.OnCollide == nil {
			continue
		}
		for _, coll2 := range *colls {
			if coll1.GameObject == coll2.GameObject {
				continue
			}
			if coll1.CanColliderWith(coll2) {
				if coll1.IsCollidedWithCollider(coll2) {
					coll1.OnCollide(coll2.GameObject)
				}
			}
		}
	}
}
