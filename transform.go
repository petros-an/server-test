package main

type Transform2D struct {
	LocalPosition Vector2D
	LocalScale    Vector2D
	localrotation Vector2D

	Parent *Transform2D
}

func (this Transform2D) Position() Vector2D {
	if this.Parent == nil {
		return this.LocalPosition
	}
	return this.LocalPosition.MulV(this.Parent.Scale()).MulC(this.Parent.Rotation()).Add(this.Parent.Position())
}

func (this Transform2D) Scale() Vector2D {
	if this.Parent == nil {
		return this.LocalScale
	}
	return this.LocalScale.MulV(this.Parent.Scale().MulConj(this.Parent.Rotation()))
}

func (this Transform2D) Rotation() Vector2D {
	if this.Parent == nil {
		return this.localrotation
	}
	return this.localrotation.MulC(this.Parent.Rotation())
}

func (this *Transform2D) SetPosition(position Vector2D) {
	if this.Parent == nil {
		this.LocalPosition = position
	}
	this.LocalPosition = position // TODO
}

func (this Transform2D) SetScale(scale Vector2D) {
	if this.Parent == nil {
		this.LocalScale = scale
	}
	this.LocalScale = scale // TODO
}

func (this Transform2D) SetRotation(rotation Vector2D) {
	if this.Parent == nil {
		this.SetLocalRotationV(rotation)
	}
	this.SetLocalRotationV(rotation) // TODO
}

func (this Transform2D) LocalRotation() Vector2D {
	return this.localrotation
}

func (this *Transform2D) SetLocalRotationV(rotation Vector2D) {
	this.localrotation = rotation.Normalized()
}

func (this *Transform2D) SetLocalRotationR(angle float64) {
	this.localrotation = newVector2DAngleR(angle)
}

func (this *Transform2D) SetLocalRotationD(angle float64) {
	this.localrotation = newVector2DAngleD(angle)
}

func (this Transform2D) Right() Vector2D {
	return this.localrotation
}

func (this Transform2D) Up() Vector2D {
	return this.localrotation.Rotate90()
}

func (this *Transform2D) SetRight(newRight Vector2D) {
	this.SetLocalRotationV(newRight)
}

func (this *Transform2D) SetUp(newUp Vector2D) {
	this.SetLocalRotationV(newUp.Rotate90Other())
}
