package geometry

import (
	"math"
)

type Circle struct {
	C Vector2D
	R float64
}

func (x *Circle) Area() float64 {
	return math.Pi * x.R * x.R
}

func (z *Circle) Copy(x *Circle) *Circle {
	z.C = x.C
	z.R = x.R
	return z
}

func (a *Circle) Equal(b *Circle) bool {
	return a.C == b.C && a.R == b.R
}

func (z *Circle) FromThreePoints(p1, p2, p3 *Vector2D) *Circle {
	return z
}

func (a *Circle) FuzzyEqual(b *Circle) bool {
	return FuzzyEqual(a.C.X, b.C.X) && FuzzyEqual(a.C.Y, b.C.Y) && FuzzyEqual(a.R, b.R)
}

func (x *Circle) Perimeter() float64 {
	return 2 * math.Pi * x.R
}
