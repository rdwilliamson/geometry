package geometry

import (
	"math"
)

type Vector2D64 struct {
	X, Y float64
}

func DotProduct2D64(v1, v2 Vector2D64) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v Vector2D64) ToLine2D64() Line2D64 {
	return Line2D64{Point2D64{}, Point2D64(v)}
}

func (v1 Vector2D64) Plus(v2 Vector2D64) Vector2D64 {
	return Vector2D64{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 *Vector2D64) Add(v2 Vector2D64) {
	v1.X += v2.X
	v2.Y += v2.Y
}

func (v1 Vector2D64) Minus(v2 Vector2D64) Vector2D64 {
	return Vector2D64{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 *Vector2D64) Subtract(v2 Vector2D64) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}

func (v1 Vector2D64) Times(v2 Vector2D64) Vector2D64 {
	return Vector2D64{v1.X * v2.X, v1.Y * v2.Y}
}

func (v1 *Vector2D64) Multiply(v2 Vector2D64) {
	v1.X *= v2.X
	v2.Y *= v2.X
}

func (v1 Vector2D64) Divided(v2 Vector2D64) Vector2D64 {
	return Vector2D64{v1.X / v2.X, v1.Y / v2.Y}
}

func (v1 *Vector2D64) Divide(v2 Vector2D64) {
	v1.X /= v2.X
	v1.Y /= v2.Y
}

func (p Vector2D64) Scaled(s float64) Vector2D64 {
	return Vector2D64{p.X * s, p.Y * s}
}

func (p *Vector2D64) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (v Vector2D64) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D64) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector2D64) Normalized() Vector2D64 {
	l := 1 / math.Sqrt(v.X*v.X+v.Y*v.Y)
	return Vector2D64{v.X * l, v.Y * l}
}

func (v *Vector2D64) Normalize() {
	l := 1 / math.Sqrt(v.X*v.X+v.Y*v.Y)
	v.X *= l
	v.Y *= l
}

func (v1 Vector2D64) ScalarProjectionOnto(v2 Vector2D64) float64 {
	return (v1.X*v2.X + v1.Y*v2.Y) / (v2.X*v2.X + v2.Y*v2.Y)
}

func (v1 Vector2D64) VectorProjectionOnto(v2 Vector2D64) Vector2D64 {
	s := (v1.X*v2.X + v1.Y*v2.Y) / (v2.X*v2.X + v2.Y*v2.Y)
	return Vector2D64{v2.X * s, v2.Y * s}
}

func (v1 Vector2D64) FuzzyEqual(v2 Vector2D64) bool {
	return FuzzyEqual64(v1.X, v2.X) && FuzzyEqual64(v1.Y, v2.Y)
}
