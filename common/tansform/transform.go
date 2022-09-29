package transform

import (
	"github.com/petros-an/server-test/common/vector"
)

type Transform2D struct {
	LocalPosition vector.Vector2D
	LocalScale    vector.Vector2D
	LocalRotation vector.Vector2D

	Parent *Transform2D
}

func (this Transform2D) Position() vector.Vector2D {
	if this.Parent == nil {
		return this.LocalPosition
	}
	return this.LocalPosition.MulV(this.Parent.Scale()).MulC(this.Parent.Rotation()).Add(this.Parent.Position())
}

func (this Transform2D) Scale() vector.Vector2D {
	if this.Parent == nil {
		return this.LocalScale
	}
	return this.LocalScale.MulV(this.Parent.Scale().MulConj(this.Parent.Rotation()))
}

func (this Transform2D) Rotation() vector.Vector2D {
	if this.Parent == nil {
		return this.LocalRotation
	}
	return this.LocalRotation.MulC(this.Parent.Rotation())
}

func (this *Transform2D) SetPosition(position vector.Vector2D) {
	if this.Parent == nil {
		this.LocalPosition = position
	}
	this.LocalPosition = position // TODO
}

func (this *Transform2D) SetScale(scale vector.Vector2D) {
	if this.Parent == nil {
		this.LocalScale = scale
	}
	this.LocalScale = scale // TODO
}

func (this *Transform2D) SetRotation(rotation vector.Vector2D) {
	if this.Parent == nil {
		this.SetLocalRotationV(rotation)
	}
	this.SetLocalRotationV(rotation) // TODO
}

// func (this Transform2D) LocalRotation() vector.Vector2D {
// 	return this.localRotation
// }

func (this *Transform2D) SetLocalRotationV(rotation vector.Vector2D) {
	this.LocalRotation = rotation.Normalized()
}

func (this *Transform2D) SetLocalRotationR(angle float64) {
	this.LocalRotation = vector.NewVector2DAngleR(angle)
}

func (this *Transform2D) SetLocalRotationD(angle float64) {
	this.LocalRotation = vector.NewVector2DAngleD(angle)
}

func (this Transform2D) Right() vector.Vector2D {
	return this.LocalRotation
}

func (this Transform2D) Up() vector.Vector2D {
	return this.LocalRotation.Rotate90()
}

func (this *Transform2D) SetRight(newRight vector.Vector2D) {
	this.SetLocalRotationV(newRight)
}

func (this *Transform2D) SetUp(newUp vector.Vector2D) {
	this.SetLocalRotationV(newUp.Rotate90Other())
}
