package geometry

import (
	"math"
)

// A Line3D represents a 3D line, line segment, or ray by a point P and vector
// V. Methods by default assume a line unless the method states otherwise.
type Line3D struct {
	P Vector3D // a point on the line
	// a vector pointing in the line or ray direction (not necessarily
	// normalized), or a vector pointing to the second end point on a line
	// segment relative to P
	V Vector3D
}

// Copy sets z to x and returns z.
func (z *Line3D) Copy(x *Line3D) *Line3D {
	z.P.X = x.P.X
	z.P.Y = x.P.Y
	z.P.Z = x.P.Z
	z.V.X = x.V.X
	z.V.Y = x.V.Y
	z.V.Z = x.V.Z
	return z
}

// Equal compares a and b then returns true if they are exactly equal or false
// otherwise.
func (a *Line3D) Equal(b *Line3D) bool {
	// check if b.P lies on a
	u := (a.V.X*(b.P.X-a.P.X) + a.V.Y*(b.P.Y-a.P.Y) + a.V.Z*(b.P.Z-a.P.Z)) / (a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z)
	if b.P.X != (a.P.X+a.V.X*u) || b.P.Y != (a.P.Y+a.V.Y*u) || b.P.Z != (a.P.Z+a.V.Z*u) {
		return false
	}
	// check if the direction of the two lines is equal
	if a.V.X == 0 && b.V.X == 0 {
		if a.V.Y == 0 && b.V.Y == 0 {
			return true
		}
		iady, ibdy := 1/a.V.Y, 1/b.V.Y
		return a.V.X*iady == b.V.X*ibdy && a.V.Z*iady == b.V.Z*ibdy
	}
	iadx, ibdx := 1/a.V.X, 1/b.V.X
	return a.V.Y*iadx == b.V.Y*ibdx && a.V.Z*iadx == b.V.Z*ibdx
}

// FuzzyEqual compares a and b and returns true if they are very close or false
// otherwise.
func (a *Line3D) FuzzyEqual(b *Line3D) bool {
	// check if b.P lies on a
	u := (a.V.X*(b.P.X-a.P.X) + a.V.Y*(b.P.Y-a.P.Y) + a.V.Z*(b.P.Z-a.P.Z)) / (a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z)
	d := math.Abs(b.P.X-(a.P.X+a.V.X*u)) + math.Abs(b.P.Y-(a.P.Y+a.V.Y*u)) + math.Abs(b.P.Z-(a.P.Z+a.V.Z*u))
	if !FuzzyEqual(d, 0) {
		return false
	}
	// check if the direction of the two lines is equal
	if FuzzyEqual(a.V.X, 0) && FuzzyEqual(b.V.X, 0) {
		if FuzzyEqual(a.V.Y, 0) && FuzzyEqual(b.V.Y, 0) {
			return true
		}
		iady, ibdy := 1/a.V.Y, 1/b.V.Y
		return FuzzyEqual(a.V.X*iady-b.V.X*ibdy, 0) && FuzzyEqual(a.V.Z*iady-b.V.Z*ibdy, 0)
	}
	iadx, ibdx := 1/a.V.X, 1/b.V.X
	return FuzzyEqual(a.V.Y*iadx-b.V.Y*ibdx, 0) && FuzzyEqual(a.V.Z*iadx-b.V.Z*ibdx, 0)
}

// Length returns the length of line segment x.
func (x *Line3D) Length() float64 {
	return math.Sqrt(x.V.X*x.V.X + x.V.Y*x.V.Y + x.V.Z*x.V.Z)
}

// LengthSquared returns the squared length of line segment x.
func (x *Line3D) LengthSquared() float64 {
	return x.V.X*x.V.X + x.V.Y*x.V.Y + x.V.Z*x.V.Z
}

// Midpoint sets point z to the line segment x's midpoint, then returns z.
func (x *Line3D) Midpoint(z *Vector3D) *Vector3D {
	z.X = x.P.X + x.V.X*0.5
	z.Y = x.P.Y + x.V.Y*0.5
	z.Z = x.P.Z + x.V.Z*0.5
	return z
}

// SegmentEqual compares line segments a and b and returns true if they are
// exactly equal or false otherwise.
func (a *Line3D) SegmentEqual(b *Line3D) bool {
	return (a.P == b.P && a.V == b.V) || (a.P.X+a.V.X == b.P.X && a.P.Y+a.V.Y == b.P.Y && a.P.Z+a.V.Z == b.P.Z &&
		-a.V.X == b.V.X && -a.V.Y == b.V.Y && -a.V.Z == b.V.Z)
}

// SegmentFuzzyEqual compares line segments a and b and returns true if they
// are very close and false otherwise.
func (a *Line3D) SegmentFuzzyEqual(b *Line3D) bool {
	return (a.P.FuzzyEqual(&b.P) && a.V.FuzzyEqual(&b.V)) || (FuzzyEqual(a.P.X+a.V.X, b.P.X) &&
		FuzzyEqual(a.P.Y+a.V.Y, b.P.Y) && FuzzyEqual(a.P.Z+a.V.Z, b.P.Z) && FuzzyEqual(-a.V.X, b.V.X) &&
		FuzzyEqual(-a.V.Y, b.V.Y) && FuzzyEqual(-a.V.Z, b.V.Z))
}
