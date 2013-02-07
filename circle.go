package geometry

import (
	"math"
)

// A Circle represents all points a given distance, R, from a point, C.
type Circle struct {
	C Vector2D
	R float64
}

// Area returns the are of the circle.
func (x *Circle) Area() float64 {
	return math.Pi * x.R * x.R
}

// Copy sets z to z then returns z.
func (z *Circle) Copy(x *Circle) *Circle {
	z.C = x.C
	z.R = x.R
	return z
}

// Equal returns true of the two circle are exactly equal or false otherwise.
func (a *Circle) Equal(b *Circle) bool {
	return a.C == b.C && a.R == b.R
}

// FromThreePoints sets z to the circle through the three points, then returns
// z.
func (z *Circle) FromThreePoints(p1, p2, p3 *Vector2D) *Circle {
	// http://paulbourke.net/geometry/circlesphere/
	if p1.X == p2.X {
		p2, p3 = p3, p2
	} else if p2.X == p3.X {
		p1, p2 = p2, p1
	}
	ma := (p2.Y - p1.Y) / (p2.X - p1.X)
	mb := (p3.Y - p2.Y) / (p3.X - p2.X)
	z.C.X = (ma*mb*(p1.Y-p3.Y) + mb*(p1.X+p2.X) - ma*(p2.X+p3.X)) / (2 * (mb - ma))
	if ma != 0 {
		z.C.Y = (-1/ma)*(z.C.X-(p1.X+p2.X)/2) + (p1.Y+p2.Y)/2
	} else {
		z.C.Y = (-1/mb)*(z.C.X-(p2.X+p3.X)/2) + (p2.Y+p3.Y)/2
	}
	dx, dy := p1.X-z.C.X, p1.Y-z.C.Y
	z.R = math.Sqrt(dx*dx + dy*dy)
	return z
}

// Fuzzy equal returns true if the two circles are very close or false
// otherwise.
func (a *Circle) FuzzyEqual(b *Circle) bool {
	return FuzzyEqual(a.C.X, b.C.X) && FuzzyEqual(a.C.Y, b.C.Y) && FuzzyEqual(a.R, b.R)
}

// Premimeter returns the perimeter of the circle.
func (x *Circle) Perimeter() float64 {
	return 2 * math.Pi * x.R
}
