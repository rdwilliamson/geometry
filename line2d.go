package geometry

import (
	"math"
)

// A Line2D represents a 2D line, line segment, or ray by a point P and vector
// V. Methods by default assume a line unless the method states otherwise.
type Line2D struct {
	P Vector2D // a point on the line
	// a vector pointing in the line or ray direction (not necessarily
	// normalized), or a vector pointing to the second end point on a line
	// segment relative to P
	V Vector2D
}

// NewLine2D returns a new Line2D.
func NewLine2D(px, py, vx, vy float64) *Line2D {
	return &Line2D{Vector2D{px, py}, Vector2D{vx, vy}}
}

// Copy sets z to x and returns z.
func (z *Line2D) Copy(x *Line2D) *Line2D {
	z.P.X = x.P.X
	z.P.Y = x.P.Y
	z.V.X = x.V.X
	z.V.Y = x.V.Y
	return z
}

// Equal compares a and b and returns true if they are exactly equal or false
// otherwise.
func (a *Line2D) Equal(b *Line2D) bool {
	am, bm := a.V.Y/a.V.X, b.V.Y/b.V.Y
	return am == bm && a.P.Y-am*a.P.X == b.P.Y-bm*b.P.X
}

// FuzzyEqual compares a and b and returns true of they are very close or false
// otherwise.
func (a *Line2D) FuzzyEqual(b *Line2D) bool {
	if FuzzyEqual(a.V.X, 0) {
		if FuzzyEqual(b.V.X, 0) {
			return FuzzyEqual(a.P.X, b.P.X)
		}
		return false
	}
	am, bm := a.V.Y/a.V.X, b.V.Y/b.V.X
	return FuzzyEqual(am, bm) && FuzzyEqual(a.P.Y-am*a.P.X, b.P.Y-bm*b.P.X)
}

// Length returns the length of line segment x.
func (x *Line2D) Length() float64 {
	return math.Sqrt(x.V.X*x.V.X + x.V.Y*x.V.Y)
}

// LengthSquared returns the squared length of line segment x.
func (x *Line2D) LengthSquared() float64 {
	return x.V.X*x.V.X + x.V.Y*x.V.Y
}

// Midpoint sets point z to the midpoint of line segment x, then returns z.
func (x *Line2D) Midpoint(z *Vector2D) *Vector2D {
	z.X = x.P.X + 0.5*x.V.X
	z.Y = x.P.Y + 0.5*x.V.Y
	return z
}

// Normal sets vector z to the normal of line x with a length equal to x's as
// if it were a line segment, then returns z.
func (x *Line2D) Normal(z *Vector2D) *Vector2D {
	z.X = x.V.Y
	z.Y = -x.V.X
	return z
}

// SegmentEqual compares a and b and returns true if the line segments are
// exactly equal and false otherwise.
func (a *Line2D) SegmentEqual(b *Line2D) bool {
	return (a.P == b.P && a.V == b.V) || (a.P == b.V && a.V == b.P)
}

// SegmentFuzzyEqual compares a and b as line segments and returns true if they
// are very close and false otherwise.
func (a *Line2D) SegmentFuzzyEqual(b *Line2D) bool {
	var ap2, bp2 Vector2D
	ap2.Add(&a.P, &a.V)
	bp2.Add(&b.P, &b.V)
	return (a.P.FuzzyEqual(&b.P) && ap2.FuzzyEqual(&bp2)) ||
		(a.P.FuzzyEqual(&bp2) && b.P.FuzzyEqual(&ap2))
}

// Slope returns the slope of x.
func (x *Line2D) Slope() float64 {
	return x.V.Y / x.V.X
}
