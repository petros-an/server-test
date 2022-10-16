package collider

import (
	"github.com/petros-an/server-test/common/collider/shape"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/game/gameObject"
)

type Collider2D interface {
	GetGameObject() gameObject.GameObject
	GetShape() shape.Shape2D

	IsCollidedWithCollider(other Collider2D) bool
	IsCollidedWithShape(shape shape.Shape2D) bool
	ShouldntExist() bool
	PrepareCollider()
	CanCollideWith(other Collider2D) bool

	HasOnCollide() bool
	OnCollide(other gameObject.GameObject)
}

type Collider2DBasic struct {
	GameObject gameObject.GameObject
	Shape      shape.Shape2D

	PositionFromTrasnform bool
	ScaleFromTrasnform    bool
	RotationFromTrasnform bool

	OnCollideField func(other gameObject.GameObject)
}

func (this *Collider2DBasic) GetGameObject() gameObject.GameObject {
	return this.GameObject
}

func (this *Collider2DBasic) GetShape() shape.Shape2D {
	return this.Shape
}

func NewBasic(gameObject gameObject.GameObject, shape shape.Shape2D) *Collider2DBasic {
	coll := Collider2DBasic{
		GameObject: gameObject,
		Shape:      shape,

		PositionFromTrasnform: true,
		ScaleFromTrasnform:    true,
		RotationFromTrasnform: true,
	}
	return &coll
}

func (this *Collider2DBasic) HasOnCollide() bool {
	return this.OnCollideField != nil
}

func (this *Collider2DBasic) OnCollide(other gameObject.GameObject) {
	this.OnCollideField(other)
}

func (this *Collider2DBasic) PrepareCollider() {

	this.Shape.SetPSR2D(this.GameObject.GetTransform().PSR2D, this.PositionFromTrasnform, this.ScaleFromTrasnform, this.RotationFromTrasnform)
	if pol, isPol := this.Shape.(*shape.Polygon); isPol {
		pol.PointsAreTransformed = false
	}
}

func (this *Collider2DBasic) IsCollidedWithCollider(other Collider2D) bool {
	return this.Shape.OverlapsShape(other.GetShape())
}

func (this *Collider2DBasic) IsCollidedWithShape(shape shape.Shape2D) bool {
	return this.Shape.OverlapsShape(shape)
}

func (this *Collider2DBasic) CanCollideWith(other Collider2D) bool {
	return true
}

func (this Collider2DBasic) ShouldntExist() bool {
	return this.GameObject.ToDestroy()
}

func CleanColliderSlice(colls *[]Collider2D) {
	len := len(*colls)
	for i := len - 1; i >= 0; i-- {
		if (*colls)[i].ShouldntExist() {
			utils.RemoveElementFromSliceAtIndex(colls, i)
		}
	}
}

func CheckCollidersCollisions(colls *[]Collider2D) {

	CleanColliderSlice(colls)

	for _, coll := range *colls {
		coll.PrepareCollider()
	}

	for _, coll1 := range *colls {
		if !coll1.HasOnCollide() {
			continue
		}
		for _, coll2 := range *colls {
			if coll1.GetGameObject() == coll2.GetGameObject() {
				continue
			}
			if coll1.CanCollideWith(coll2) {
				if coll1.IsCollidedWithCollider(coll2) {
					coll1.OnCollide(coll2.GetGameObject())
				}
			}
		}
	}
}
