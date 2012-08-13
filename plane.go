package geometry

import (
	"math"
)

// A Plane represents a plane by ax + by + cz + d = 0. A, B, and C represent
// the normal direction only (could have any magnitude). D divided by the
// magnitude of the normal represents the distance of the plane from the
// origin and the sign, if positive, signifies the origin is in the half space
// determined by the direction of the normal or the other half space if
// negative.
type Plane struct {
	A, B, C, D float64
}

// NewPLane returns a new Plane.
func NewPlane(a, b, c, d float64) *Plane {
	return &Plane{a, b, c, d}
}

// Equal returns true if the two planes are exactly equal or false otherwise.
func (a *Plane) Equal(b *Plane) bool {
	// check normal and distance from origin direction
	s := a.A / b.A
	if a.B != s*b.B || a.C != s*b.C || s*a.D*b.D < 0 {
		return false
	}
	// check distance (squared) from origin
	return b.D*b.D*(a.A*a.A+a.B*a.B+a.C*a.C) ==
		a.D*a.D*(b.A*b.A+b.B*b.B+b.C*b.C)
}

// FuzzyEqual returns true if the two planes are very close or false otherwise.
func (a *Plane) FuzzyEqual(b *Plane) bool {
	s := a.A / b.A
	if s*a.D*b.D < 0 || !FuzzyEqual(a.B, s*b.B) || !FuzzyEqual(a.C, s*b.C) {
		return false
	}
	return FuzzyEqual(b.D*b.D*(a.A*a.A+a.B*a.B+a.C*a.C),
		a.D*a.D*(b.A*b.A+b.B*b.B+b.C*b.C))
}

// LineIntersection sets z to the intersecion of plane a and line b, then
// returns z.
func (a *Plane) LineIntersection(b *Line3D, z *Vector3D) *Vector3D {
	// http://paulbourke.net/geometry/planeline/
	bdx, bdy, bdz := b.P1.X-b.P2.X, b.P1.Y-b.P2.Y, b.P1.Z-b.P2.Z
	u := (a.A*b.P1.X + a.B*b.P1.Y + a.C*b.P1.Z + a.D) /
		(a.A*bdx + a.B*bdy + a.C*bdz)
	z.X = b.P1.X - u*bdx
	z.Y = b.P1.Y - u*bdy
	z.Z = b.P1.Z - u*bdz
	return z
}

// Nomrmal sets z to x's normal then returns z.
func (x *Plane) Normal(z *Vector3D) *Vector3D {
	z.X = x.A
	z.Y = x.B
	z.Z = x.C
	return z
}

// Normalize sets z to the Hessian normal form of x where the normal is a unit
// vector and D is the distance from the origin, then returns z.
func (x *Plane) Normalize(z *Plane) *Plane {
	s := 1 / math.Sqrt(x.A*x.A+x.B*x.B+x.C*x.C)
	z.A = x.A * s
	z.B = x.B * s
	z.C = x.C * s
	z.D = x.D * s
	return z
}

// NormalizedEqual returns true if the two planes (assumed to be in Hessian
// normal form) are exactly equal or false otherwise.
func (a *Plane) NormalizedEqual(b *Plane) bool {
	if a.D*b.D < 0 {
		return -a.A == b.A && -a.B == b.B && -a.C == b.C && -a.D == b.D
	}
	return a.A == b.A && a.B == b.B && a.C == b.C && a.D == b.D
}

// PointDistance returns the distance (may be negative) point b is from plane a
// assuming it is normalized.
func (a *Plane) NormalizedPointDistance(b *Vector3D) float64 {
	return a.A*b.X + a.B*b.Y + a.C*b.Z + a.D
}

// PointDistance returns the distance (may be negative) point b is from plane a.
func (a *Plane) PointDistance(b *Vector3D) float64 {
	return (a.A*b.X + a.B*b.Y + a.C*b.Z + a.D) /
		math.Sqrt(a.A*a.A+a.B*a.B+a.C*a.C)
}

// PointDistance returns the squared distance point b is from plane a.
func (a *Plane) PointDistanceSquared(b *Vector3D) float64 {
	n := a.A*b.X + a.B*b.Y + a.C*b.Z + a.D
	return (n * n) / (a.A*a.A + a.B*a.B + a.C*a.C)
}

// Set sets z to x then returns z.
func (z *Plane) Set(x *Plane) *Plane {
	z.A = x.A
	z.B = x.B
	z.C = x.C
	z.D = x.D
	return z
}

// SetFromPoints set z to the plane through the three points, then returns z.
func (z *Plane) SetFromPoints(p1, p2, p3 *Vector3D) *Plane {
	// http://paulbourke.net/geometry/planeeq/
	z.A = p1.Y*(p2.Z-p3.Z) + p2.Y*(p3.Z-p1.Z) + p3.Y*(p1.Z-p2.Z)
	z.B = p1.Z*(p2.X-p3.X) + p2.Z*(p3.X-p1.X) + p3.Z*(p1.X-p2.X)
	z.C = p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y)
	z.D = -(p1.X*(p2.Y*p3.Z-p3.Y*p2.Z) + p2.X*(p3.Y*p1.Z-p1.Y*p3.Z) +
		p3.X*(p1.Y*p2.Z-p2.Y*p1.Z))
	return z
}

// ThreePlaneIntersection sets z to the intersection of planes a, b, and c,
// then returns z.
func (a *Plane) ThreePlaneIntersection(b, c *Plane, z *Vector3D) *Vector3D {
	// http://paulbourke.net/geometry/3planes/
	return z
}

// TwoPlaneIntersection sets z to the intersection of planes a and b, then
// returns z.
func (a *Plane) TwoPlaneIntersection(b *Plane, z *Line3D) *Line3D {
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
