package geometry

import (
	"math"
)

// A line represented by two points.
type Line3D struct {
	P1, P2 Point3D
}

// Returns a vector from the first point to the second.
func (l *Line3D) ToVector() Vector3D {
	return Vector3D{l.P2.X - l.P1.X, l.P2.Y - l.P1.Y, l.P2.Z - l.P1.Z}
}

// Returns the length of the vector.
func (l *Line3D) Length() float64 {
	dx, dy, dz := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y, l.P2.Z-l.P1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Returns the length of the vector.
func (l *Line3D) LengthSquared() float64 {
	dx, dy, dz := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y, l.P2.Z-l.P1.Z
	return dx*dx + dy*dy + dz*dz
}

// Returns the midpoint of the line.
func (l *Line3D) Midpoint() Point3D {
	return Point3D{(l.P2.X - l.P1.X) * 0.5, (l.P2.Y - l.P1.Y) * 0.5, (l.P2.Z - l.P1.Z) * 0.5}
}

// Returns true if the lines are equal.
func (l1 *Line3D) Equal(l2 *Line3D) bool {
	return (l1.P1 == l2.P1 && l1.P2 == l2.P2) || (l1.P1 == l2.P2 && l1.P2 == l2.P1)
}

// Returns true if the lines are nearly equal.
func (l1 *Line3D) FuzzyEqual(l2 *Line3D) bool {
	dx1, dy1, dz1 := l1.P1.X-l2.P1.X, l1.P1.Y-l2.P1.Y, l1.P1.Z-l2.P1.Z
	dx2, dy2, dz2 := l1.P2.X-l2.P2.X, l1.P2.Y-l2.P2.Y, l1.P2.Z-l2.P2.Z
	if dx1*dx1+dy1*dy1+dz1*dz1 < 0.000000000001*0.000000000001 &&
		dx2*dx2+dy2*dy2+dz2*dz2 < 0.000000000001*0.000000000001 {
		return true
	}
	dx1, dy1, dz1 = l1.P1.X-l2.P2.X, l1.P1.Y-l2.P2.Y, l1.P1.Z-l2.P2.Z
	dx2, dy2, dz2 = l1.P2.X-l2.P1.X, l1.P2.Y-l2.P1.Y, l1.P2.Z-l2.P1.Z
	return dx1*dx1+dy1*dy1+dz1*dz1 < 0.000000000001*0.000000000001 &&
		dx2*dx2+dy2*dy2+dz2*dz2 < 0.000000000001*0.000000000001
}

// Returns the distance between a point and a line.
func (l *Line3D) PointDistance(p *Point3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y, l.P2.Z-l.P1.Z
	u := (ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y) + ldz*(p.Z-l.P1.Z)) / (ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := p.X-(l.P1.X+ldx*u), p.Y-(l.P1.Y+ldy*u), p.Z-(l.P1.Z+ldz*u)
	return math.Sqrt(x*x + y*y + z*z)
}

// Returns the squared distance between a point and a line.
func (l *Line3D) PointSquaredDistance(p *Point3D) float64 {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/pointline/
	ldx, ldy, ldz := l.P2.X-l.P1.X, l.P2.Y-l.P1.Y, l.P2.Z-l.P1.Z
	u := (ldx*(p.X-l.P1.X) + ldy*(p.Y-l.P1.Y) + ldz*(p.Z-l.P1.Z)) / (ldx*ldx + ldy*ldy + ldz*ldz)
	x, y, z := p.X-(l.P1.X+ldx*u), p.Y-(l.P1.Y+ldy*u), p.Z-(l.P1.Z+ldz*u)
	return x*x + y*y + z*z
}

// Returns the shortest line between two lines.
func (l1 *Line3D) LineBetween(l2 *Line3D) Line3D {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline3d/
	l1dx, l1dy, l1dz := l1.P2.X-l1.P1.X, l1.P2.Y-l1.P1.Y, l1.P2.Z-l1.P1.Z
	l2dx, l2dy, l2dz := l2.P2.X-l2.P1.X, l2.P2.Y-l2.P1.Y, l2.P2.Z-l2.P1.Z
	p1dx, p1dy, p1dz := l1.P1.X-l2.P1.X, l1.P1.Y-l2.P1.Y, l1.P1.Z-l2.P1.Z
	d1343 := p1dx*l2dx + p1dy*l2dy + p1dz*l2dz
	d4321 := l2dx*l1dx + l2dy*l1dy + l2dz*l1dz
	d1321 := p1dx*l1dx + p1dy*l1dy + p1dz*l1dz
	d4343 := l2dx*l2dx + l2dy*l2dy + l2dz*l2dz
	d2121 := l1dx*l1dx + l1dy*l1dy + l1dz*l2dz
	mua := (d1343*d4321 - d1321*d4343) / (d2121*d4343 - d4321*d4321)
	mub := (d1343 + mua*d4321) / d4343
	return Line3D{Point3D{l1dx*mua + l1.P1.X, l1dy*mua + l1.P1.Y, l1dz*mua + l1.P1.Z},
		Point3D{l2dx*mub + l2.P1.X, l2dy*mub + l2.P1.Y, l2dz*mub + l2.P1.Z}}
}
