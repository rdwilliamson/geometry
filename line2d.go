package geometry

import (
	"math"
)

type Line2D struct {
	P1, P2 Point2D
}

// http://softsurfer.com/Archive/algorithm_0102/algorithm_0102.htm
func LinePointDistance2D64(l Line2D, p Point2D, segment bool) float64 {
	return 0
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
