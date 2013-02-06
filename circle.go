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
	// from paulbourke.net/geometry/circlesphere
	if p1.X == p2.X {
		tmp := p2
		p2 = p3
		p3 = tmp
	} else if p2.X == p3.X {
		tmp := p2
		p2 = p1
		p1 = tmp
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

func (a *Circle) FuzzyEqual(b *Circle) bool {
	return FuzzyEqual(a.C.X, b.C.X) && FuzzyEqual(a.C.Y, b.C.Y) && FuzzyEqual(a.R, b.R)
}

func (x *Circle) Perimeter() float64 {
	return 2 * math.Pi * x.R
}
