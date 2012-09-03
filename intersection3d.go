package geometry

import (
	"math"
)

// Intersection3DLineLine, sets z to the shortest line between a and b then
// returns 1.
func Intersection3DLineLine(a, b, z *Line3D) int {
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
	return 1
}

// Intersection3DLineSegmentLineSegment determines the shortest line segment
// between a and b, then returns 1.
func Intersection3DLineSegmentLineSegment(a, b, z *Line3D) int {
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
	return 1
}

// Intersection3DPlaneLine sets z to the intersecion of plane a and line b,
// then returns 1.
func Intersection3DPlaneLine(a *Plane, b *Line3D, z *Vector3D) int {
	// http://paulbourke.net/geometry/planeline/
	bdx, bdy, bdz := b.P1.X-b.P2.X, b.P1.Y-b.P2.Y, b.P1.Z-b.P2.Z
	u := (a.A*b.P1.X + a.B*b.P1.Y + a.C*b.P1.Z + a.D) /
		(a.A*bdx + a.B*bdy + a.C*bdz)
	z.X = b.P1.X - u*bdx
	z.Y = b.P1.Y - u*bdy
	z.Z = b.P1.Z - u*bdz
	return 1
}

// Intersection3DPlanePlane sets z to the intersection of planes a and b, then
// returns 1.
func Intersection3DPlanePlane(a, b *Plane, z *Line3D) int {
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
	return 1
}

// Intersection3DPlanePlanePlane sets z to the intersection of planes a, b, and
// c, then returns 1.
func Intersection3DPlanePlanePlane(a, b, c *Plane, z *Vector3D) int {
	// http://paulbourke.net/geometry/3planes/
	n2n3x := b.B*c.C - b.C*c.B
	n2n3y := b.C*c.A - b.A*c.C
	n2n3z := b.A*c.B - b.B*c.A
	n3n1x := c.B*a.C - c.C*a.B
	n3n1y := c.C*a.A - c.A*a.C
	n3n1z := c.A*a.B - c.B*a.A
	n1n2x := a.B*b.C - a.C*b.B
	n1n2y := a.C*b.A - a.A*b.C
	n1n2z := a.A*b.B - a.B*b.A
	d := -1 / (a.A*n2n3x + a.B*n2n3y + a.C*n2n3z)
	z.X = (a.D*n2n3x + b.D*n3n1x + c.D*n1n2x) * d
	z.Y = (a.D*n2n3y + b.D*n3n1y + c.D*n1n2y) * d
	z.Z = (a.D*n2n3z + b.D*n3n1z + c.D*n1n2z) * d
	return 1
}

// Intersection3DFuzzyPlaneLine sets z to the intersection of a plane and line,
// then returns the number of intersections.
//
// Possible return values are:
// -1 if the plane and line are coincident, z in untouched.
// 0 if the plane and line are parallel, z is untouched.
// 1 if an intersection occurs, z is set to the intersection point.
func Intersection3DFuzzyPlaneLine(a *Plane, b *Line3D, z *Vector3D) int {
	// http://paulbourke.net/geometry/planeline/
	bdx, bdy, bdz := b.P1.X-b.P2.X, b.P1.Y-b.P2.Y, b.P1.Z-b.P2.Z
	dot2 := a.A*bdx + a.B*bdy + a.C*bdz
	dot1 := a.A*b.P1.X + a.B*b.P1.Y + a.C*b.P1.Z
	if FuzzyEqual(dot2, 0) {
		if FuzzyEqual(dot1, 0) {
			return -1
		}
		return 0
	}
	u := (dot1 + a.D) / dot2
	z.X = b.P1.X - u*bdx
	z.Y = b.P1.Y - u*bdy
	z.Z = b.P1.Z - u*bdz
	return 1
}

// Intersection3DFuzzyPlanePlane sets z to the intersection of planes a and b,
// then returns the number of intersections.
//
// Possible return values are:
// -1 if the planes are coincident, z is untouched.
// 0 if the planes are parallel, z is untouched.
// 1 if an intersection occurs, z is set to the intersection line.
func Intersection3DFuzzyPlanePlane(a, b *Plane, z *Line3D) int {
	// http://paulbourke.net/geometry/planeplane/
	cpx, cpy, cpz := a.B*b.C-a.C*b.B, a.C*b.A-a.A*b.C, a.A*b.B-a.B*b.A
	n1n1 := a.A*a.A + a.B*a.B + a.C*a.C
	n2n2 := b.A*b.A + b.B*b.B + b.C*b.C
	if FuzzyEqual(cpx*cpx+cpy*cpy+cpz*cpz, 0) {
		// TODO a.A or b.A almost zero
		s := a.A / b.A
		if s*a.D*b.D < 0 || !FuzzyEqual(a.B, s*b.B) ||
			!FuzzyEqual(a.C, s*b.C) ||
			!FuzzyEqual(b.D*b.D*n1n1, a.D*a.D*n2n2) {
			return 0
		}
		return -1
	}
	n1n2 := a.A*b.A + a.B*b.B + a.C*b.C
	d := 1 / (n1n1*n2n2 - n1n2*n1n2)
	c1 := (b.D*n1n2 - a.D*n2n2) * d
	c2 := (a.D*n1n2 - b.D*n1n1) * d
	z.P1.X = c1*a.A + c2*b.A
	z.P1.Y = c1*a.B + c2*b.B
	z.P1.Z = c1*a.C + c2*b.C
	z.P2.X = z.P1.X + cpx
	z.P2.Y = z.P1.Y + cpy
	z.P2.Z = z.P1.Z + cpz
	return 1
}

// Intersection3DFuzzyPlanePlanePlane sets z to the intersecion of 3 planes,
// then returns the number of intersections.
//
// Possible return values are:
// -3 if two planes are parallel and the third intersects at two lines.
// -2 if all three planes intersect at a line.
// -1 if all three planes are coincident.
// 0 if all planes are parallel, z is untouched.
// 1 if the planes intersect at a point, z is set to the intersection point.
func Intersection3DFuzzyPlanePlanePlane(a, b, c *Plane, z *Vector3D) int {
	var n1, n2, n3 Vector3D
	a.Normal(&n1)
	b.Normal(&n2)
	c.Normal(&n3)
	if vectorDirectionEqual(&n1, &n2) && vectorDirectionEqual(&n1, &n3) {
		if FuzzyEqual(b.D*b.D*(a.A*a.A+a.B*a.B+a.C*a.C),
			a.D*a.D*(b.A*b.A+b.B*b.B+b.C*b.C)) &&
			FuzzyEqual(c.D*c.D*(a.A*a.A+a.B*a.B+a.C*a.C),
				a.D*a.D*(c.A*c.A+c.B*c.B+c.C*c.C)) {
			return -1
		}
	}
	var l1d, l2d Vector3D
	l1d.CrossProduct(&n1, &n2)
	l2d.CrossProduct(&n2, &n3)
	if vectorDirectionEqual(&l1d, &l2d) {
		n1n1 := n1.DotProduct(&n1)
		n1n2 := n1.DotProduct(&n2)
		n2n2 := n2.DotProduct(&n2)
		n2n3 := n2.DotProduct(&n3)
		n3n3 := n3.DotProduct(&n3)
		d1 := 1 / (n1n1*n2n2 - n1n2*n1n2)
		d2 := 1 / (n2n2*n3n3 - n2n3*n2n3)
		c1 := (b.D*n1n2 - a.D*n2n2) * d1
		c2 := (a.D*n1n2 - b.D*n1n1) * d1
		c3 := (c.D*n2n3 - b.D*n3n3) * d2
		c4 := (b.D*n2n3 - c.D*n2n2) * d2
		p1 := Vector3D{c1*a.A + c2*b.A, c1*a.B + c2*b.B, c1*a.C + c2*b.C}
		p2 := Vector3D{c3*b.A + c4*c.A, c3*b.B + c4*c.B, c3*b.C + c4*c.C}
		if FuzzyEqual(p1.X*c.A+p1.Y*c.B+p1.Z*c.C+c.D, 0) &&
			FuzzyEqual(p2.X*a.A+p2.Y*a.B+p2.Z*a.C+a.D, 0) {
			return -2
		}
	}
	if n1.FuzzyEqual(&n2) && n2.FuzzyEqual(&n3) {
		return 0
	}
	if vectorDirectionEqual(&n1, &n2) || vectorDirectionEqual(&n2, &n3) ||
		vectorDirectionEqual(&n1, &n3) {
		return -3
	}
	Intersection3DPlanePlanePlane(a, b, c, z)
	return 1
}

func vectorDirectionEqual(a, b *Vector3D) bool {
	if FuzzyEqual(math.Abs(a.X)+math.Abs(b.X), 0) {
		if FuzzyEqual(math.Abs(a.Y)+math.Abs(b.Y), 0) {
			return true
		} else {
			if a.Y > b.Y {
				s := math.Abs(a.Y / b.Y)
				return FuzzyEqual(a.Z, s*b.Z)
			}
			s := math.Abs(b.Y / a.Y)
			return FuzzyEqual(s*a.Z, b.Z)
		}
	}
	if a.X > b.X {
		s := math.Abs(a.X / b.X)
		return FuzzyEqual(a.Y, s*b.Y) && FuzzyEqual(a.Z, s*b.Z)
	}
	s := math.Abs(b.X / a.X)
	return FuzzyEqual(s*a.Y, b.Y) && FuzzyEqual(s*a.Z, b.Z)
}
