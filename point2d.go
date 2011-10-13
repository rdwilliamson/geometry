package geometry

import (
	"fmt"
	"math"
)

type Point2D struct {
	X, Y float64
}

func (p1 Point2D) DistTo(p2 Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p1 Point2D) DistToSq(p2 Point2D) float64 {
	dx, dy := p2.X-p1.X, p2.Y-p1.Y
	return dx*dx + dy*dy
}

func (p Point2D) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point2D) LengthSq() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Point2D) Normalized() Point2D {
	s := 1.0 / math.Sqrt(p.X*p.X+p.Y*p.Y)
	return Point2D{p.X * s, p.Y * s}
}

func (p1 Point2D) Plus(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y}
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%g, %g)", p.X, p.Y)
}

func (p1 Point2D) Minus(p2 Point2D) Point2D {
	return Point2D{p1.X - p2.X, p1.Y - p2.Y}
}

func (p Point2D) Scaled(s float64) Point2D {
	return Point2D{p.X * s, p.Y * s}
}

func (p Point2D) ToPoint3D(z float64) Point3D {
	return Point3D{p.X, p.Y, z}
}

func DotProduct2D(p1, p2 Point2D) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}
