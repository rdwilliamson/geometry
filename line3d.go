package geometry

import (
	"math"
)

// A Line3D representes a 3D line by two points P1 and P2 (represented by
// vectors) on the line. The line is treated as an infinite line unless a
// method explicitly says otherwise. If treated as a segment then P1 and P2 are
// the end points of the line segment.
type Line3D struct {
	P1, P2 Vector3D
}

// Copy sets z to x and returns z.
func (z *Line3D) Copy(x *Line3D) *Line3D {
	z.P1.X = x.P1.X
	z.P1.Y = x.P1.Y
	z.P1.Z = x.P1.Z
	z.P2.X = x.P2.X
	z.P2.Y = x.P2.Y
	z.P2.Z = x.P2.Z
	return z
}

// Equal compares a and b then returns true if they are exactly equal or false
// otherwise.
func (a *Line3D) Equal(b *Line3D) bool {
	// check if b.P1 lies on a
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (adx*(b.P1.X-a.P1.X) + ady*(b.P1.Y-a.P1.Y) + adz*(b.P1.Z-a.P1.Z)) /
		(adx*adx + ady*ady + adz*adz)
	if b.P1.X != (a.P1.X+adx*u) || b.P1.Y != (a.P1.Y+adx*u) ||
		b.P1.Z != (a.P1.Z+adx*u) {
		return false
	}
	// check if the direction of the two lines is equal
	iadx, ibdx := 1/adx, 1/(b.P2.X-b.P1.X)
	return ady*iadx == (b.P2.Y-b.P1.Y)*ibdx &&
		adz*iadx == (b.P2.Z-b.P1.Z)*ibdx
}

// FromPlanesIntersection sets z to the intersection of planes a and b, then
// returns z.
func (z *Line3D) FromPlanesIntersection(a, b *Plane) *Line3D {
	// http://paulbourke.net/geometry/planeplane/
	n1n1 := a.A*a.A + a.B*a.B + a.C*a.C
	n2n2 := b.A*b.A + b.B*b.B + b.C*b.C
	n1n2 := a.A*b.A + a.B*b.B + a.C*b.C
	d := 1 / (n1n1*n2n2 - n1n2*n1n2)
	c1 := (b.D*n1n2 - a.D*n2n2) * d
	c2 := (a.D*n1n2 - b.D*n1n1) * d
	z.P1.X = c1*a.A + c2*b.A
	z.P1.Y = c1*a.B + c2*b.B
	z.P1.Z = c1*a.C + c2*b.C
	z.P2.X = z.P1.X + (a.B*b.C - a.C*b.B)
	z.P2.Y = z.P1.Y + (a.C*b.A - a.A*b.C)
	z.P2.Z = z.P1.Z + (a.A*b.B - a.B*b.A)
	return z
}

// FuzzyEqual compares a and b and returns true if they are very close or false
// otherwise.
func (a *Line3D) FuzzyEqual(b *Line3D) bool {
	// check if b.P1 lies on a
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (adx*(b.P1.X-a.P1.X) + ady*(b.P1.Y-a.P1.Y) + adz*(b.P1.Z-a.P1.Z)) /
		(adx*adx + ady*ady + adz*adz)
	d := math.Abs(b.P1.X - (a.P1.X + adx*u))
	d += math.Abs(b.P1.Y - (a.P1.Y + ady*u))
	d += math.Abs(b.P1.Z - (a.P1.Z + adz*u))
	if !FuzzyEqual(d, 0) {
		return false
	}
	// check if the direction of the two lines is equal
	iadx, ibdx := 1/adx, 1/(b.P2.X-b.P1.X)
	dyr := ady*iadx - (b.P2.Y-b.P1.Y)*ibdx
	dzr := adz*iadx - (b.P2.Z-b.P1.Z)*ibdx
	return FuzzyEqual(dyr, 0) && FuzzyEqual(dzr, 0)
}

// Length returns the length of line segment x.
func (x *Line3D) Length() float64 {
	dx, dy, dz := x.P2.X-x.P1.X, x.P2.Y-x.P1.Y, x.P2.Z-x.P1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// LengthSquared returns the squared length of line segment x.
func (x *Line3D) LengthSquared() float64 {
	dx, dy, dz := x.P2.X-x.P1.X, x.P2.Y-x.P1.Y, x.P2.Z-x.P1.Z
	return dx*dx + dy*dy + dz*dz
}

// Midpoint sets point z to the line segment x's midpoint, then returns z.
func (x *Line3D) Midpoint(z *Vector3D) *Vector3D {
	z.X = (x.P1.X + x.P2.X) * 0.5
	z.Y = (x.P1.Y + x.P2.Y) * 0.5
	z.Z = (x.P1.Z + x.P2.Z) * 0.5
	return z
}

// SegmentEqual compares line segments a and b and returns true if they are
// exactly equal or false otherwise.
func (a *Line3D) SegmentEqual(b *Line3D) bool {
	return (a.P1 == b.P1 && a.P2 == b.P2) || (a.P1 == b.P2 && a.P2 == b.P1)
}

// SegmentFuzzyEqual compares line segments a and b and returns true if they
// are very close and false otherwise.
func (a *Line3D) SegmentFuzzyEqual(b *Line3D) bool {
	return (a.P1.FuzzyEqual(&b.P1) && a.P2.FuzzyEqual(&b.P2)) ||
		(a.P1.FuzzyEqual(&b.P2) && a.P2.FuzzyEqual(&b.P1))
}

// ToVector sets z to the vector from a.P1 to a.P2, then returns z.
func (x *Line3D) ToVector(z *Vector3D) *Vector3D {
	z.X = x.P2.X - x.P1.X
	z.Y = x.P2.Y - x.P1.Y
	z.Z = x.P2.Z - x.P1.Z
	return z
}
