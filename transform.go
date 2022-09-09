package main

type Transform struct {
	Position Vector2D
	Scale    Vector2D
	rotation Vector2D

	Parent *Transform
}

func (this Transform) FinalPosition() Vector2D {
	if this.Parent == nil {
		return this.Position
	}
	return this.Position.MulV(this.Parent.FinalScale()).MulC(this.Parent.FinalRotation()).Add(this.Parent.FinalPosition())
}

func (this Transform) FinalScale() Vector2D {
	if this.Parent == nil {
		return this.Scale
	}
	return this.Scale.MulV(this.Parent.FinalScale().MulConj(this.Parent.FinalRotation()))
}

func (this Transform) FinalRotation() Vector2D {
	if this.Parent == nil {
		return this.rotation
	}
	return this.rotation.MulC(this.Parent.FinalRotation())
}

func (this Transform) Rotation() Vector2D {
	return this.rotation
}

func (this *Transform) SetRotationV(rotation Vector2D) {
	this.rotation = rotation.Normalized()
}

func (this *Transform) SetRotationR(angle float64) {
	this.rotation = newVector2DAngleR(angle)
}

func (this *Transform) SetRotationD(angle float64) {
	this.rotation = newVector2DAngleD(angle)
}

func (this Transform) Right() Vector2D {
	return this.rotation
}

func (this Transform) Up() Vector2D {
	return this.rotation.Rotate90()
}

func (this *Transform) SetRight(newRight Vector2D) {
	this.SetRotationV(newRight)
}

func (this *Transform) SetUp(newUp Vector2D) {
	this.SetRotationV(newUp.Rotate90Other())
}
