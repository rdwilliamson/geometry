package geometry

import (
	"fmt"
	"math"
)

type Line2D struct {
	P1, P2 Point2D
}

func (l *Line2D) Angle() float64 {
	dx, dy := l.P2.X - l.P1.X, l.P2.Y - l.P1.Y
	return math.Atan(dy/dx)
}

func (l *Line2D) Length() float64 {
	dx, dy := l.P2.X - l.P1.X, l.P2.Y - l.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (l Line2D) String() string {
	return fmt.Sprintf("{%v %v}", l.P1, l.P2)
}
