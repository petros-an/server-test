package main

import (
	"math"
)

type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (this Vector2D) Magnitude() float64 {
	return math.Sqrt(this.X*this.X + this.Y + this.Y)
}

func (this Vector2D) MagnitudeSq() float64 {
	return this.X*this.X + this.Y + this.Y
}

func (this Vector2D) AngleR() float64 {
	return math.Atan2(this.Y, this.X)
}

func (this Vector2D) AngleD() float64 {
	return math.Atan2(this.Y, this.X) * 180 / math.Pi
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
	this.X = this.X*other.X - this.Y*other.Y
	this.Y = this.X*other.Y + this.Y*other.X
	return this
}

func (this Vector2D) MulConj(other Vector2D) Vector2D {
	return Vector2D{this.X*other.X + this.Y*other.Y, this.X*other.Y - this.Y*other.X}
}

func (this *Vector2D) MulConjSelf(other Vector2D) *Vector2D {
	this.X = this.X*other.X + this.Y*other.Y
	this.Y = this.X*other.Y - this.Y*other.X
	return this
}

func (this Vector2D) DivC(other Vector2D) Vector2D {
	return Vector2D{this.X*other.X + this.Y*other.Y, this.X*other.Y - this.Y*other.X}.Div(other.MagnitudeSq())
}

func (this *Vector2D) DivCSelf(other Vector2D) *Vector2D {
	this.X = this.X*other.X + this.Y*other.Y
	this.Y = this.X*other.Y - this.Y*other.X
	return this.DivSelf(other.MagnitudeSq())
}
