package geometry

import (
	"fmt"
	"math"
)

type Point3D struct {
	X, Y, Z float64
}

func (p Point3D) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p Point3D) LengthSquared() float64 {
	return p.X*p.X + p.Y*p.Y + p.Z*p.Z
}

func (p Point3D) Normalized() Point3D {
	s := 1.0 / math.Sqrt(p.X*p.X+p.Y*p.Y+p.Z*p.Z)
	return Point3D{p.X * s, p.Y * s, p.Z * s}
}

func (p Point3D) Scaled(s float64) Point3D {
	return Point3D{p.X * s, p.Y * s, p.Z * s}
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%g, %g, %g)", p.X, p.Y, p.Z)
}

func (p Point3D) ToPoint2D() Point2D {
	return Point2D{p.X, p.Y}
}
