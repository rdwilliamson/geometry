package geometry

import (
	"math"
)

// A Line2D representes a 2D line by two points P1 and P2 (represented by
// vectors) on the line. The line is treated as an infinite line unless a
// method explicitly says otherwise. If treated as a segment then P1 and P2 are
// the end points of the line segment.
type Line2D struct {
	P1, P2 Vector2D
}

// NewLine2D returns a new Line2D.
func NewLine2D(x1, y1, x2, y2 float64) *Line2D {
	return &Line2D{Vector2D{x1, y1}, Vector2D{x2, y2}}
}

// Should rays have P1 be the end point and P2 treated as a vector or P2 as a
// point on the ray? I never use rays so I'm not sure which is more convenient.

// Copy sets z to x and returns z.
func (z *Line2D) Copy(x *Line2D) *Line2D {
	z.P1.X = x.P1.X
	z.P1.Y = x.P1.Y
	z.P2.X = x.P2.X
	z.P2.Y = x.P2.Y
	return z
}

// Equal compares a and b and returns true if they are exactly equal or false
// otherwise.
func (a *Line2D) Equal(b *Line2D) bool {
	am, bm := (a.P2.Y-a.P1.Y)/(a.P2.X-a.P1.X), (b.P2.Y-b.P1.Y)/(b.P2.X-b.P1.X)
	return am == bm && a.P1.Y-am*a.P1.X == b.P1.Y-bm*b.P1.X
}

// FuzzyEqual compares a and b and returns true of they are very close or false
// otherwise.
func (a *Line2D) FuzzyEqual(b *Line2D) bool {
	am, bm := (a.P2.Y-a.P1.Y)/(a.P2.X-a.P1.X), (b.P2.Y-b.P1.Y)/(b.P2.X-b.P1.X)
	return FuzzyEqual(am, bm) && FuzzyEqual(a.P1.Y-am*a.P1.X, b.P1.Y-bm*b.P1.X)
}

// Length returns the length of line segment x.
func (x *Line2D) Length() float64 {
	dx, dy := x.P2.X-x.P1.X, x.P2.Y-x.P1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// LengthSquared returns the squared length of line segment x.
func (x *Line2D) LengthSquared() float64 {
	dx, dy := x.P2.X-x.P1.X, x.P2.Y-x.P1.Y
	return dx*dx + dy*dy
}

// Midpoint sets point z to the midpoint of line segment x, then returns z.
func (x *Line2D) Midpoint(z *Vector2D) *Vector2D {
	z.X = (x.P1.X + x.P2.X) * 0.5
	z.Y = (x.P1.Y + x.P2.Y) * 0.5
	return z
}

// Normal sets vector z to the normal of line x with a length equal to x's as
// if it were a line segment, then returns z.
func (x *Line2D) Normal(z *Vector2D) *Vector2D {
	z.X = x.P2.Y - x.P1.Y
	z.Y = x.P1.X - x.P2.X
	return z
}

// SegmentEqual compares a and b and returns true if the line segments are
// exactly equal and false otherwise.
func (a *Line2D) SegmentEqual(b *Line2D) bool {
	return (a.P1 == b.P1 && a.P2 == b.P2) || (a.P1 == b.P2 && a.P2 == b.P1)
}

// SegmentFuzzyEqual compares a and b as line segments and returns true if they
// are very close and false otherwise.
func (a *Line2D) SegmentFuzzyEqual(b *Line2D) bool {
	return (a.P1.FuzzyEqual(&b.P1) && a.P2.FuzzyEqual(&b.P2)) ||
		(a.P1.FuzzyEqual(&b.P2) && a.P2.FuzzyEqual(&b.P1))
}

// Slope returns the slope of x.
func (x *Line2D) Slope() float64 {
	return (x.P2.Y - x.P1.Y) / (x.P2.X - x.P1.X)
}

// ToVector sets z to the vector from x.P1 to x.P2, then returns z.
func (x *Line2D) ToVector(z *Vector2D) *Vector2D {
	z.X = x.P2.X - x.P1.X
	z.Y = x.P2.Y - x.P1.Y
	return z
}
