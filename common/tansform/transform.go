package transform

import (
	"github.com/petros-an/server-test/common/vector"
)

type Transform2D struct {
	vector.PSR2D

	Parent *Transform2D
}

func New(
	position vector.Vector2D,
	scale vector.Vector2D,
	rotation vector.Vector2D,
) Transform2D {
	return Transform2D{
		PSR2D: vector.PSR2D{
			Position: position,
			Scale:    scale,
			Rotation: rotation,
		},
	}
}

func (this Transform2D) FinalPosition() vector.Vector2D {
	if this.Parent == nil {
		return this.Position
	}
	return this.Position.MulV(this.Parent.Scale).MulC(this.Parent.FinalRotation()).Add(this.Parent.FinalPosition())
}

func (this Transform2D) FinalRotation() vector.Vector2D {
	if this.Parent == nil {
		return this.Rotation
	}
	return this.Rotation.MulC(this.Parent.FinalRotation())
}

func (this *Transform2D) SetPosition(position vector.Vector2D) {
	if this.Parent == nil {
		this.Position = position
	}
	this.Position = position // TODO
}

func (this *Transform2D) SetScale(scale vector.Vector2D) {
	if this.Parent == nil {
		this.Scale = scale
	}
	this.Scale = scale // TODO
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
	this.Rotation = rotation.Normalized()
}

func (this *Transform2D) SetLocalRotationR(angle float64) {
	this.Rotation = vector.NewVector2DAngleR(angle)
}

func (this *Transform2D) SetLocalRotationD(angle float64) {
	this.Rotation = vector.NewVector2DAngleD(angle)
}

func (this Transform2D) Right() vector.Vector2D {
	return this.Rotation
}

func (this Transform2D) Up() vector.Vector2D {
	return this.Rotation.Rotate90()
}

func (this *Transform2D) SetRight(newRight vector.Vector2D) {
	this.SetLocalRotationV(newRight)
}

func (this *Transform2D) SetUp(newUp vector.Vector2D) {
	this.SetLocalRotationV(newUp.Rotate90Other())
}
