package geometry

import (
	"math"
)

type Vector2D struct {
	X, Y float64
}

func DotProduct2D(v1, v2 Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v Vector2D) ToLine2D() Line2D {
	return Line2D{Point2D{0, 0}, Point2D(v)}
}

func (v1 Vector2D) Plus(v2 Vector2D) Vector2D {
	return Vector2D{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 *Vector2D) Add(v2 Vector2D) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v1 Vector2D) Minus(v2 Vector2D) Vector2D {
	return Vector2D{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 *Vector2D) Subtract(v2 Vector2D) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}

func (v1 Vector2D) Times(v2 Vector2D) Vector2D {
	return Vector2D{v1.X * v2.X, v1.Y * v2.Y}
}

func (v1 *Vector2D) Multiply(v2 Vector2D) {
	v1.X *= v2.X
	v1.Y *= v2.X
}

func (v1 Vector2D) Divided(v2 Vector2D) Vector2D {
	return Vector2D{v1.X / v2.X, v1.Y / v2.Y}
}

func (v1 *Vector2D) Divide(v2 Vector2D) {
	v1.X /= v2.X
	v1.Y /= v2.Y
}

func (p Vector2D) Scaled(s float64) Vector2D {
	return Vector2D{p.X * s, p.Y * s}
}

func (p *Vector2D) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (v Vector2D) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vector2D) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector2D) Normalized() Vector2D {
	l := 1 / math.Hypot(v.X, v.Y)
	return Vector2D{v.X * l, v.Y * l}
}

func (v *Vector2D) Normalize() {
	l := 1 / math.Hypot(v.X, v.Y)
	v.X *= l
	v.Y *= l
}

func (v1 Vector2D) ScalarProjectionOnto(v2 Vector2D) float64 {
	return (v1.X*v2.X + v1.Y*v2.Y) / (v2.X*v2.X + v2.Y*v2.Y)
}

func (v1 Vector2D) VectorProjectionOnto(v2 Vector2D) Vector2D {
	s := (v1.X*v2.X + v1.Y*v2.Y) / (v2.X*v2.X + v2.Y*v2.Y)
	return Vector2D{v2.X * s, v2.Y * s}
}

func (v1 Vector2D) Equal(v2 Vector2D) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 Vector2D) FuzzyEqual(v2 Vector2D) bool {
	return FuzzyEqual(v1.X, v2.X) && FuzzyEqual(v1.Y, v2.Y)
}
