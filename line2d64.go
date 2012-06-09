package geometry

import (
	"math"
)

type Line2D64 struct {
	P1, P2 Point2D64
}

// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func LinePointDistance2D64(l Line2D64, p Point2D64, segment bool) float64 {
	v := l.ToVector2D64()
	w := Vector2D64(p).Minus(Vector2D64(l.P1))
	c1 := DotProduct2D64(w, v)
	if segment && c1 <= 0 {
		return Vector2D64(l.P1.Minus(p)).Length()
	}
	c2 := DotProduct2D64(v, v)
	if segment && c2 <= c1 {
		return Vector2D64(l.P2.Minus(p)).Length()
	}
	b := c1 / c2
	a := l.P1.Plus(Point2D64(v.Scaled(b)))
	return Vector2D64(a.Minus(p)).Length()
}

func (l Line2D64) ToVector2D64() Vector2D64 {
	return Vector2D64{l.P2.X - l.P1.X, l.P2.Y - l.P1.X}
}

func (l Line2D64) Dx() float64 {
	return l.P2.X - l.P1.X
}

func (l Line2D64) Dy() float64 {
	return l.P2.Y - l.P1.Y
}

func (l Line2D64) Length() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (l Line2D64) LengthSquared() float64 {
	dx := l.P2.X - l.P1.X
	dy := l.P2.Y - l.P1.Y
	return dx*dx + dy*dy
}

func (l Line2D64) Midpoint() Point2D64 {
	return Point2D64{(l.P1.X + l.P2.X) * 0.5, (l.P1.Y + l.P2.Y) * 0.5}
}

func (l Line2D64) Angle() float64 {
	return math.Atan((l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X))
}
