package geometry

import (
	"math"
)

type Line2D struct {
	P1, P2 Point2D
}

// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func LinePointDistance2D64(l Line2D, p Point2D, segment bool) float64 {
	v := l.ToVector2D()
	w := Vector2D(p).Minus(Vector2D(l.P1))
	c1 := DotProduct2D(w, v)
	if segment && c1 <= 0 {
		return Vector2D(l.P1.Minus(p)).Length()
	}
	c2 := DotProduct2D(v, v)
	if segment && c2 <= c1 {
		return Vector2D(l.P2.Minus(p)).Length()
	}
	b := c1 / c2
	a := l.P1.Plus(Point2D(v.Scaled(b)))
	return Vector2D(a.Minus(p)).Length()
}

func (l Line2D) ToVector2D() Vector2D {
	return Vector2D{l.P2.X - l.P1.X, l.P2.Y - l.P1.X}
}

func (l Line2D) Dx() float64 {
	return l.P2.X - l.P1.X
}

func (l Line2D) Dy() float64 {
	return l.P2.Y - l.P1.Y
}

func (l Line2D) Length() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (l Line2D) LengthSquared() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return dx*dx + dy*dy
}

func (l Line2D) Midpoint() Point2D {
	return Point2D{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

func (l Line2D) Angle() float64 {
	return math.Atan((l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X))
}
