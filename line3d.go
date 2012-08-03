package geometry

import "math"

// A Line3D representes a 3D line by two points P1 and P2 (represented by
// vectors) on the line. The line is treated as an infinite line unless a
// method explicitly says otherwise. If treated as a segment then P1 and P2 are
// the end points of the line segment.
type Line3D struct {
	P1, P2 Vector3D
}

// Equal compares a and b and returns a boolean indicating if they are equal.
func (a *Line3D) Equal(b *Line3D) bool {
	l1dx, l1dy, l1dz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	d := 1 / (l1dx*l1dx + l1dy*l1dy + l1dz*l1dz)
	u := (l1dx*(b.P1.X-a.P1.X) + l1dy*(b.P1.Y-a.P1.Y) + l1dz*(b.P1.Z-a.P1.Z)) * d
	x, y, z := b.P1.X-(a.P1.X+l1dx*u), b.P1.Y-(a.P1.Y+l1dy*u), b.P1.Z-(a.P1.Z+l1dz*u)
	if x*x+y*y+z*z != 0 {
		return false
	}
	il1dx, il2dx := 1/l1dx, 1/(b.P2.X-b.P1.X)
	return l1dy*il1dx == (b.P2.Y-b.P1.Y)*il2dx &&
		l1dz*il1dx == (b.P2.Z-b.P1.Z)*il2dx
}

// FuzzyEqual compares a and b and returns a boolean indicating if they are
// very close.
func (a *Line3D) FuzzyEqual(b *Line3D) bool {
	l1dx, l1dy, l1dz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	d := 1 / (l1dx*l1dx + l1dy*l1dy + l1dz*l1dz)
	u := (l1dx*(b.P1.X-a.P1.X) + l1dy*(b.P1.Y-a.P1.Y) + l1dz*(b.P1.Z-a.P1.Z)) * d
	x, y, z := b.P1.X-(a.P1.X+l1dx*u), b.P1.Y-(a.P1.Y+l1dy*u), b.P1.Z-(a.P1.Z+l1dz*u)
	if x*x+y*y+z*z >= 0.000000000001*0.000000000001 {
		return false
	}
	il1dx, il2dx := 1/l1dx, 1/(b.P2.X-b.P1.X)
	dyr := l1dy*il1dx - (b.P2.Y-b.P1.Y)*il2dx
	dzr := l1dz*il1dx - (b.P2.Z-b.P1.Z)*il2dx
	return dyr*dyr < 0.000000000001*0.000000000001 &&
		dzr*dzr < 0.000000000001*0.000000000001
}

// LineBetween sets z to the shortest line between a and b and returns z. This
// function is intended as a replacement for intersection (which can be still
// be tested by z.P1 == z.P2).
func (z *Line3D) LineBetween(a, b *Line3D) *Line3D {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline3d/
	adx, ady, adz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	bdx, bdy, bdz := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y, b.P2.Z-b.P1.Z
	p1dx, p1dy, p1dz := a.P1.X-b.P1.X, a.P1.Y-b.P1.Y, a.P1.Z-b.P1.Z
	d1343 := p1dx*bdx + p1dy*bdy + p1dz*bdz
	d4321 := bdx*adx + bdy*ady + bdz*adz
	d1321 := p1dx*adx + p1dy*ady + p1dz*adz
	d4343 := bdx*bdx + bdy*bdy + bdz*bdz
	d2121 := adx*adx + ady*ady + adz*bdz
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

// Length returns the length of l as if is a line segment.
func (a *Line3D) Length() float64 {
	dx, dy, dz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// LengthSquared returns the length squared of l as if is a line segment.
func (a *Line3D) LengthSquared() float64 {
	dx, dy, dz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	return dx*dx + dy*dy + dz*dz
}

// Midpoint sets z to the segment l's midpoint and returns z.
func (a *Line3D) Midpoint(z *Vector3D) *Vector3D {
	z.X = (a.P1.X + a.P2.X) * 0.5
	z.Y = (a.P1.Y + a.P2.Y) * 0.5
	z.Z = (a.P1.Z + a.P2.Z) * 0.5
	return z
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
func (a *Line3D) PointSquaredDistance(b *Vector3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y, a.P2.Z-a.P1.Z
	u := (ldx*(b.X-a.P1.X) + ldy*(b.Y-a.P1.Y) + ldz*(b.Z-a.P1.Z)) /
		(ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := b.X-(a.P1.X+ldx*u), b.Y-(a.P1.Y+ldy*u), b.Z-(a.P1.Z+ldz*u)
	return x*x + y*y + z*z
}

// SegmentEqual compares a and b as line segments and returns a boolean
// indicating if they are equal.
func (a *Line3D) SegmentEqual(b *Line3D) bool {
	return (a.P1 == b.P1 && a.P2 == b.P2) || (a.P1 == b.P2 && a.P2 == b.P1)
}

// SegmentFuzzyEqual compares a and b as line segments and returns a boolean
// indicating if they are very close.
func (a *Line3D) SegmentFuzzyEqual(b *Line3D) bool {
	dx1, dy1, dz1 := a.P1.X-b.P1.X, a.P1.Y-b.P1.Y, a.P1.Z-b.P1.Z
	dx2, dy2, dz2 := a.P2.X-b.P2.X, a.P2.Y-b.P2.Y, a.P2.Z-b.P2.Z
	if dx1*dx1+dy1*dy1+dz1*dz1 < 0.000000000001*0.000000000001 &&
		dx2*dx2+dy2*dy2+dz2*dz2 < 0.000000000001*0.000000000001 {
		return true
	}
	dx1, dy1, dz1 = a.P1.X-b.P2.X, a.P1.Y-b.P2.Y, a.P1.Z-b.P2.Z
	dx2, dy2, dz2 = a.P2.X-b.P1.X, a.P2.Y-b.P1.Y, a.P2.Z-b.P1.Z
	return dx1*dx1+dy1*dy1+dz1*dz1 < 0.000000000001*0.000000000001 &&
		dx2*dx2+dy2*dy2+dz2*dz2 < 0.000000000001*0.000000000001
}

// SegmentIntersection sets z to the shortest line between a and b and returns
// a boolean indicating if both z's end points lie on line segments a and b.
// This function is intended as a replacement for intersection (which can be
// still be tested by z.P1 == z.P2).

// SegmentPointDistance returns the distance between line segment a and point
// b.

// SegmentPointDistanceSquared returns the squared distance between line
// segment a and point b.

// Set sets z to x and returns z.
func (z *Line3D) Set(x *Line3D) *Line3D {
	z.P1.X = x.P1.X
	z.P1.Y = x.P1.Y
	z.P1.Z = x.P1.Z
	z.P2.X = x.P2.X
	z.P2.Y = x.P2.Y
	z.P2.Z = x.P2.Z
	return z
}

// ToVector sets z to the vector from l.P1 to l.P2 and returns z.
func (x *Line3D) ToVector(z *Vector3D) *Vector3D {
	z.X = x.P2.X - x.P1.X
	z.Y = x.P2.Y - x.P1.Y
	z.Z = x.P2.Z - x.P1.Z
	return z
}
