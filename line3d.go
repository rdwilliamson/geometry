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

// LineBetween sets z to the shortest line between a and b then returns z. This
// function is intended as a replacement for intersection (which can be still
// be tested by z.P1 == z.P2).
func (a *Line3D) LineBetween(b, z *Line3D) *Line3D {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline3d/
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	bdx, bdy, bdz := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y, b.P2.Z-b.P1.Z
	p1dx, p1dy, p1dz := a.P1.X-b.P1.X, a.P1.Y-b.P1.Y, a.P1.Z-b.P1.Z
	d1343 := p1dx*bdx + p1dy*bdy + p1dz*bdz
	d4321 := bdx*adx + bdy*ady + bdz*adz
	d1321 := p1dx*adx + p1dy*ady + p1dz*adz
	d4343 := bdx*bdx + bdy*bdy + bdz*bdz
	d2121 := adx*adx + ady*ady + adz*adz
	mua := (d1343*d4321 - d1321*d4343) / (d2121*d4343 - d4321*d4321)
	mub := (d1343 + mua*d4321) / d4343
	z.P1.X = adx*mua + a.P1.X
	z.P1.Y = ady*mua + a.P1.Y
	z.P1.Z = adz*mua + a.P1.Z
	z.P2.X = bdx*mub + b.P1.X
	z.P2.Y = bdy*mub + b.P1.Y
	z.P2.Z = bdz*mub + b.P1.Z
	return z
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

// PointAngularDistance returns the angle the line segment a would have to
// rotate about its midpoint to pass through point b.
func (a *Line3D) PointAngularDistance(b *Vector3D) float64 {
	mpx, mpy, mpz := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5, (a.P1.Z+a.P2.Z)*0.5
	l1dx, l1dy, l1dz := a.P1.X-mpx, a.P1.Y-mpy, a.P1.Z-mpz
	l2dx, l2dy, l2dz := b.X-mpx, b.Y-mpy, b.Z-mpz
	return math.Abs(math.Acos((l1dx*l2dx+l1dy*l2dy+l1dz*l2dz)/
		math.Sqrt((l1dx*l1dx+l1dy*l1dy+l1dz*l1dz)*(l2dx*l2dx+l2dy*l2dy+l2dz*l2dz))) - math.Pi/2)
}

// PointAngularDistanceCosSquared returns the cos of the squared angle the line
// segment a would have to rotate about its midpoint to pass through point b.
func (a *Line3D) PointAngularDistanceCosSquared(b *Vector3D) float64 {
	mpx, mpy, mpz := (a.P1.X+a.P2.X)*0.5, (a.P1.Y+a.P2.Y)*0.5, (a.P1.Z+a.P2.Z)*0.5
	l1dx, l1dy, l1dz := a.P1.X-mpx, a.P1.Y-mpy, a.P1.Z-mpz
	l2dx, l2dy, l2dz := b.X-mpx, b.Y-mpy, b.Z-mpz
	dot := l1dx*l2dx + l1dy*l2dy + l1dz*l2dz
	return dot * dot / ((l1dx*l1dx + l1dy*l1dy + l1dz*l1dz) * (l2dx*l2dx + l2dy*l2dy + l2dz*l2dz))
}

// PointDistance returns the distance point b is from line a.
func (a *Line3D) PointDistance(b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y) + ldz*(b.Z-a.P1.Z)) /
		(ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u), b.Z-(a.P1.Z+ldz*u)
	return math.Sqrt(x*x + y*y + z*z)
}

// PointDistanceSquared returns the squared distance point b is from line a.
func (a *Line3D) PointDistanceSquared(b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y) + ldz*(b.Z-a.P1.Z)) /
		(ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u), b.Z-(a.P1.Z+ldz*u)
	return x*x + y*y + z*z
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

// SegmentLineBetween sets z to the shortest line segment between a and b then
// returns z. This function is intended as a replacement for intersection
// (which can be still be tested by z.P1 == z.P2).
func (a *Line3D) SegmentLineBetween(b, z *Line3D) *Line3D {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline3d/
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	bdx, bdy, bdz := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y, b.P2.Z-b.P1.Z
	p1dx, p1dy, p1dz := a.P1.X-b.P1.X, a.P1.Y-b.P1.Y, a.P1.Z-b.P1.Z
	d1343 := p1dx*bdx + p1dy*bdy + p1dz*bdz
	d4321 := bdx*adx + bdy*ady + bdz*adz
	d1321 := p1dx*adx + p1dy*ady + p1dz*adz
	d4343 := bdx*bdx + bdy*bdy + bdz*bdz
	d2121 := adx*adx + ady*ady + adz*adz
	mua := (d1343*d4321 - d1321*d4343) / (d2121*d4343 - d4321*d4321)
	mub := (d1343 + mua*d4321) / d4343
	if mua < 0 {
		z.P1.X = a.P1.X
		z.P1.Y = a.P1.Y
		z.P1.Z = a.P1.Z
	} else if mua > 1 {
		z.P1.X = a.P2.X
		z.P1.Y = a.P2.Y
		z.P1.Z = a.P2.Z
	} else {
		z.P1.X = adx*mua + a.P1.X
		z.P1.Y = ady*mua + a.P1.Y
		z.P1.Z = adz*mua + a.P1.Z
	}
	if mub < 0 {
		z.P2.X = b.P1.X
		z.P2.Y = b.P1.Y
		z.P2.Z = b.P1.Z
	} else if mub > 1 {
		z.P2.X = b.P2.X
		z.P2.Y = b.P2.Y
		z.P2.Z = b.P2.Z
	} else {
		z.P2.X = bdx*mub + b.P1.X
		z.P2.Y = bdy*mub + b.P1.Y
		z.P2.Z = bdz*mub + b.P1.Z
	}
	return z
}

// SegmentPointDistance returns the distance between line segment a and point
// b.
func (a *Line3D) SegmentPointDistance(b *Vector3D) float64 {
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	b1dx, b1dy, b1dz := b.X-a.P1.X, b.Y-a.P1.Y, b.Z-a.P1.Z
	u := (adx*b1dx + ady*b1dy + adz*b1dz) / (adx*adx + ady*ady + adz*adz)
	var x, y, z float64
	if u < 0 {
		x, y, z = b1dx, b1dy, b1dz
	} else if u > 1 {
		x, y, z = b.X-a.P2.X, b.Y-a.P2.Y, b.Z-a.P2.Z
	} else {
		x, y, z = b.X-(a.P1.X+adx*u), b.Y-(a.P1.Y+ady*u), b.Z-(a.P1.Z+adz*u)
	}
	return math.Sqrt(x*x + y*y + z*z)
}

// SegmentPointDistanceSquared returns the squared distance between line
// segment a and point b.
func (a *Line3D) SegmentPointDistanceSquared(b *Vector3D) float64 {
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	b1dx, b1dy, b1dz := b.X-a.P1.X, b.Y-a.P1.Y, b.Z-a.P1.Z
	u := (adx*b1dx + ady*b1dy + adz*b1dz) / (adx*adx + ady*ady + adz*adz)
	var x, y, z float64
	if u < 0 {
		x, y, z = b1dx, b1dy, b1dz
	} else if u > 1 {
		x, y, z = b.X-a.P2.X, b.Y-a.P2.Y, b.Z-a.P2.Z
	} else {
		x, y, z = b.X-(a.P1.X+adx*u), b.Y-(a.P1.Y+ady*u), b.Z-(a.P1.Z+adz*u)
	}
	return x*x + y*y + z*z
}

// ToVector sets z to the vector from l.P1 to l.P2 and returns z.
func (x *Line3D) ToVector(z *Vector3D) *Vector3D {
	z.X = x.P2.X - x.P1.X
	z.Y = x.P2.Y - x.P1.Y
	z.Z = x.P2.Z - x.P1.Z
	return z
}
