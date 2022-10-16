package vector

import (
	"fmt"
	"math"
	"math/rand"
)

type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (this Vector2D) String() string {
	return fmt.Sprintf("(%f, %f)", this.X, this.Y)
}

func New(x float64, y float64) Vector2D {
	return Vector2D{X: x, Y: y}
}

func RandomNew() Vector2D {
	return Vector2D{X: rand.Float64()*80 - 40, Y: rand.Float64()*80 - 40}
}

func Zero() Vector2D {
	return Vector2D{X: 0, Y: 0}
}

func NewVector2DAngleR(angle float64) Vector2D {
	return Vector2D{X: math.Cos(angle), Y: math.Sin(angle)}
}
func NewVector2DMagnAngleR(magnitude float64, angle float64) Vector2D {
	return Vector2D{X: magnitude * math.Cos(angle), Y: magnitude * math.Sin(angle)}
}

func NewVector2DAngleD(angle float64) Vector2D {
	return NewVector2DAngleR(angle * math.Pi / 180)
}
func NewVector2DMagnAngleD(magnitude float64, angle float64) Vector2D {
	return NewVector2DMagnAngleR(magnitude, angle*math.Pi/180)
}

func (this Vector2D) Magnitude() float64 {
	m := this.X*this.X + this.Y*this.Y
	if m == 1 {
		return 1
	}
	return math.Sqrt(m)
}

func (this *Vector2D) SetMagnitude(newMagnitude float64) *Vector2D {
	return this.Normalize().MulSelf(newMagnitude)
}

func (this Vector2D) MagnitudeSq() float64 {
	return this.X*this.X + this.Y*this.Y
}

func (this Vector2D) MagnitudeManhattan() float64 {
	return math.Abs(this.X+this.Y) + math.Abs(this.X-this.Y)
}

func (this Vector2D) AngleR() float64 {
	return math.Atan2(this.Y, this.X)
}

func (this Vector2D) AngleD() float64 {
	return math.Atan2(this.Y, this.X) * 180 / math.Pi
}

func (this Vector2D) Normalized() Vector2D {
	m := this.MagnitudeSq()
	if m == 1 {
		return this
	}
	m = math.Sqrt(m)
	return Vector2D{X: this.X / m, Y: this.Y / m}
}

func (this *Vector2D) Normalize() *Vector2D {
	m := this.MagnitudeSq()
	if m == 1 {
		return this
	}
	m = math.Sqrt(m)
	this.X /= m
	this.Y /= m
	return this
}

func (this Vector2D) Conj() Vector2D {
	return Vector2D{X: this.X, Y: -this.Y}
}

func (this *Vector2D) ConjSelf() *Vector2D {
	this.Y = -this.Y
	return this
}

func (this Vector2D) AbsValues() Vector2D {
	return Vector2D{X: math.Abs(this.X), Y: math.Abs(this.Y)}
}

func (this *Vector2D) AbsValuesSelf() *Vector2D {
	this.X = math.Abs(this.X)
	this.Y = math.Abs(this.Y)
	return this
}

func (this Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{this.X + other.X, this.Y + other.Y}
}

func (this *Vector2D) AddSelf(other Vector2D) *Vector2D {
	this.X += other.X
	this.Y += other.Y
	return this
}

func (this Vector2D) Sub(other Vector2D) Vector2D {
	return Vector2D{this.X - other.X, this.Y - other.Y}
}

func (this *Vector2D) SubSelf(other Vector2D) *Vector2D {
	this.X -= other.X
	this.Y -= other.Y
	return this
}

func (this Vector2D) Opp() Vector2D {
	return Vector2D{X: -this.X, Y: -this.Y}
}

func (this *Vector2D) OppSelf() *Vector2D {
	this.X, this.Y = -this.X, -this.Y
	return this
}

func (this Vector2D) Mul(a float64) Vector2D {
	return Vector2D{this.X * a, this.Y * a}
}

func (this *Vector2D) MulSelf(a float64) *Vector2D {
	this.X *= a
	this.Y *= a
	return this
}

func (this Vector2D) Div(a float64) Vector2D {
	return Vector2D{this.X / a, this.Y / a}
}

func (this *Vector2D) DivSelf(a float64) *Vector2D {
	this.X /= a
	this.Y /= a
	return this
}

func (this Vector2D) MulV(other Vector2D) Vector2D {
	return Vector2D{this.X * other.X, this.Y * other.Y}
}

func (this *Vector2D) MulVSelf(other Vector2D) *Vector2D {
	this.X *= other.X
	this.Y *= other.Y
	return this
}

func (this Vector2D) DivV(other Vector2D) Vector2D {
	return Vector2D{this.X / other.X, this.Y / other.Y}
}

func (this *Vector2D) DivVSelf(other Vector2D) *Vector2D {
	this.X /= other.X
	this.Y /= other.Y
	return this
}

func (this Vector2D) MulC(other Vector2D) Vector2D {
	return Vector2D{this.X*other.X - this.Y*other.Y, this.X*other.Y + this.Y*other.X}
}

func (this *Vector2D) MulCSelf(other Vector2D) *Vector2D {
	this.X, this.Y = this.X*other.X-this.Y*other.Y, this.X*other.Y+this.Y*other.X
	return this
}

func (this Vector2D) MulConj(other Vector2D) Vector2D {
	return Vector2D{this.X*other.X + this.Y*other.Y, this.X*other.Y - this.Y*other.X}
}

func (this *Vector2D) MulConjSelf(other Vector2D) *Vector2D {
	this.X, this.Y = this.X*other.X+this.Y*other.Y, this.X*other.Y-this.Y*other.X
	return this
}

func (this Vector2D) DivC(other Vector2D) Vector2D {
	return Vector2D{this.X*other.X + this.Y*other.Y, this.X*other.Y - this.Y*other.X}.Div(other.MagnitudeSq())
}

func (this *Vector2D) DivCSelf(other Vector2D) *Vector2D {
	this.X, this.Y = this.X*other.X+this.Y*other.Y, this.X*other.Y-this.Y*other.X
	return this.DivSelf(other.MagnitudeSq())
}

func (this Vector2D) Rotate90() Vector2D {
	return Vector2D{X: -this.Y, Y: this.X}
}

func (this *Vector2D) Rotate90Self() *Vector2D {
	this.X, this.Y = -this.Y, this.X
	return this
}

func (this Vector2D) Rotate90Other() Vector2D {
	return Vector2D{X: this.Y, Y: -this.X}
}

func (this *Vector2D) Rotate90OtherSelf() *Vector2D {
	this.X, this.Y = this.Y, -this.X
	return this
}

func (this Vector2D) RotateAroundCenter(rotation Vector2D, center Vector2D) Vector2D {
	new := this
	return *new.RotateAroundCenterSelf(rotation, center)
}

func (this *Vector2D) RotateAroundCenterSelf(rotation Vector2D, center Vector2D) *Vector2D {
	return this.SubSelf(center).MulCSelf(rotation).AddSelf(center)
}

func (this Vector2D) ScaleFromCenter(scale Vector2D, center Vector2D) Vector2D {
	new := this
	return *new.ScaleFromCenterSelf(scale, center)
}

func (this *Vector2D) ScaleFromCenterSelf(scale Vector2D, center Vector2D) *Vector2D {
	return this.SubSelf(center).MulVSelf(scale).AddSelf(center)
}

func (this Vector2D) UnscaleFromCenter(scale Vector2D, center Vector2D) Vector2D {
	new := this
	return *new.UnScaleFromCenterSelf(scale, center)
}

func (this *Vector2D) UnScaleFromCenterSelf(scale Vector2D, center Vector2D) *Vector2D {
	return this.SubSelf(center).DivVSelf(scale).AddSelf(center)
}

func (this Vector2D) InvertXY() Vector2D {
	return Vector2D{X: this.Y, Y: this.X}
}

func (this *Vector2D) InvertXYSelf() *Vector2D {
	this.X, this.Y = this.Y, this.X
	return this
}

func (this Vector2D) Equals(other Vector2D) bool {
	return this.X == other.X && this.Y == other.Y
}

func (this Vector2D) Dot(other Vector2D) float64 {
	return this.X*other.X + this.Y*other.Y
}

func (this Vector2D) DotNum(x float64, y float64) float64 {
	return this.X*x + this.Y*y
}

type PSR2D struct {
	Position Vector2D
	Scale    Vector2D
	Rotation Vector2D
}

func (this *PSR2D) SetPSR2D(other PSR2D, setPosition bool, setScale bool, setRotation bool) {
	if setPosition {
		this.Position = other.Position
	}
	if setScale {
		this.Scale = other.Scale
	}
	if setRotation {
		this.Rotation = other.Rotation
	}
}

func (this PSR2D) RotationDifference(other PSR2D) Vector2D {
	if this.Rotation.Equals(other.Rotation) {
		return Vector2D{X: 1, Y: 0}
	}
	return this.Rotation.MulConj(other.Rotation)
}

func (this PSR2D) HasUniformScale() bool {
	return this.Scale.X == this.Scale.Y
}

func (this Vector2D) ApplyTransformation(psr PSR2D) Vector2D {
	new := this
	return *new.ApplyTransformationSelf(psr)
}

func (this *Vector2D) ApplyTransformationSelf(psr PSR2D) *Vector2D {
	this.MulVSelf(psr.Scale)
	if psr.Rotation.Y != 0 {
		this.MulCSelf(psr.Rotation)
	}
	return this.AddSelf(psr.Position)
}

func (this Vector2D) RemoveTransformation(psr PSR2D) Vector2D {
	new := this
	return *new.RemoveTransformationSelf(psr)
}

func (this *Vector2D) RemoveTransformationSelf(psr PSR2D) *Vector2D {
	this.SubSelf(psr.Position)
	if psr.Rotation.Y != 0 {
		return this.MulConjSelf(psr.Rotation)
	}
	return this.DivVSelf(psr.Scale)
}
