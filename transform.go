package main

type Transform2D struct {
	Position Vector2D
	Scale    Vector2D
	rotation Vector2D

	Parent *Transform2D
}

func (this Transform2D) FinalPosition() Vector2D {
	if this.Parent == nil {
		return this.Position
	}
	return this.Position.MulV(this.Parent.FinalScale()).MulC(this.Parent.FinalRotation()).Add(this.Parent.FinalPosition())
}

func (this Transform2D) FinalScale() Vector2D {
	if this.Parent == nil {
		return this.Scale
	}
	return this.Scale.MulV(this.Parent.FinalScale().MulConj(this.Parent.FinalRotation()))
}

func (this Transform2D) FinalRotation() Vector2D {
	if this.Parent == nil {
		return this.rotation
	}
	return this.rotation.MulC(this.Parent.FinalRotation())
}

func (this Transform2D) Rotation() Vector2D {
	return this.rotation
}

func (this *Transform2D) SetRotationV(rotation Vector2D) {
	this.rotation = rotation.Normalized()
}

func (this *Transform2D) SetRotationR(angle float64) {
	this.rotation = newVector2DAngleR(angle)
}

func (this *Transform2D) SetRotationD(angle float64) {
	this.rotation = newVector2DAngleD(angle)
}

func (this Transform2D) Right() Vector2D {
	return this.rotation
}

func (this Transform2D) Up() Vector2D {
	return this.rotation.Rotate90()
}

func (this *Transform2D) SetRight(newRight Vector2D) {
	this.SetRotationV(newRight)
}

func (this *Transform2D) SetUp(newUp Vector2D) {
	this.SetRotationV(newUp.Rotate90Other())
}
