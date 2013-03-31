package geometry

import (
	"math"
)

// Intersection3DLineLine, sets z to the shortest line between a and b then
// returns 1.
func Intersection3DLineLine(a, b, z *Line3D) int {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline3d/
	pdx, pdy, pdz := a.P.X-b.P.X, a.P.Y-b.P.Y, a.P.Z-b.P.Z
	d1343 := pdx*b.V.X + pdy*b.V.Y + pdz*b.V.Z
	d4321 := b.V.X*a.V.X + b.V.Y*a.V.Y + b.V.Z*a.V.Z
	d1321 := pdx*a.V.X + pdy*a.V.Y + pdz*a.V.Z
	d4343 := b.V.X*b.V.X + b.V.Y*b.V.Y + b.V.Z*b.V.Z
	d2121 := a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z
	mua := (d1343*d4321 - d1321*d4343) / (d2121*d4343 - d4321*d4321)
	mub := (d1343 + mua*d4321) / d4343
	z.P.X = a.V.X*mua + a.P.X
	z.P.Y = a.V.Y*mua + a.P.Y
	z.P.Z = a.V.Z*mua + a.P.Z
	z.V.X = (b.V.X*mub + b.P.X) - z.P.X
	z.V.Y = (b.V.Y*mub + b.P.Y) - z.P.Y
	z.V.Z = (b.V.Z*mub + b.P.Z) - z.P.Z
	return 1
}

func Intersection3DLineSphere(a *Line3D, b *Sphere, y, z *Vector3D) int {
	// http://paulbourke.net/geometry/circlesphere/index.html
	aa := a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z
	bb := 2 * (a.V.X*(a.P.X-b.C.X) + a.V.Y*(a.P.Y-b.C.Y) + a.V.Z*(a.P.Z-b.C.Z))
	cc := b.C.X*b.C.X + b.C.Y*b.C.Y + b.C.Z*b.C.Z + a.P.X*a.P.X + a.P.Y*a.P.Y + a.P.Z*a.P.Z
	cc -= 2 * (b.C.X*a.P.X + b.C.Y*a.P.Y + b.C.Z*a.P.Z)
	cc -= b.R * b.R
	rr := bb*bb - 4*aa*cc
	if rr < 0 {
		return 0
	}
	if rr == 0 {
		u := -bb / (2 * aa)
		y.X = a.P.X + u*a.V.X
		y.Y = a.P.Y + u*a.V.Y
		y.Z = a.P.Z + u*a.V.Z
		return 1
	}
	sr := math.Sqrt(rr)
	aa = 1 / (2 * aa)
	u := (-bb + sr) * aa
	y.X = a.P.X + u*a.V.X
	y.Y = a.P.Y + u*a.V.Y
	y.Z = a.P.Z + u*a.V.Z
	u = (-bb - sr) * aa
	z.X = a.P.X + u*a.V.X
	z.Y = a.P.Y + u*a.V.Y
	z.Z = a.P.Z + u*a.V.Z
	return 2
}

// Intersection3DLineSegmentLineSegment determines the shortest line segment
// between a and b, then returns 1.
func Intersection3DLineSegmentLineSegment(a, b, z *Line3D) int {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline3d/
	pdx, pdy, pdz := a.P.X-b.P.X, a.P.Y-b.P.Y, a.P.Z-b.P.Z
	d1343 := pdx*b.V.X + pdy*b.V.Y + pdz*b.V.Z
	d4321 := b.V.X*a.V.X + b.V.Y*a.V.Y + b.V.Z*a.V.Z
	d1321 := pdx*a.V.X + pdy*a.V.Y + pdz*a.V.Z
	d4343 := b.V.X*b.V.X + b.V.Y*b.V.Y + b.V.Z*b.V.Z
	d2121 := a.V.X*a.V.X + a.V.Y*a.V.Y + a.V.Z*a.V.Z
	mua := (d1343*d4321 - d1321*d4343) / (d2121*d4343 - d4321*d4321)
	mub := (d1343 + mua*d4321) / d4343
	if mua < 0 {
		z.P.X = a.P.X
		z.P.Y = a.P.Y
		z.P.Z = a.P.Z
	} else if mua > 1 {
		z.P.X = a.P.X + a.V.X
		z.P.Y = a.P.Y + a.V.Y
		z.P.Z = a.P.Z + a.V.Z
	} else {
		z.P.X = a.V.X*mua + a.P.X
		z.P.Y = a.V.Y*mua + a.P.Y
		z.P.Z = a.V.Z*mua + a.P.Z
	}
	if mub < 0 {
		z.V.X = b.P.X - z.P.X
		z.V.Y = b.P.Y - z.P.Y
		z.V.Z = b.P.Z - z.P.Z
	} else if mub > 1 {
		z.V.X = (b.P.X + b.V.X) - z.P.X
		z.V.Y = (b.P.Y + b.V.Y) - z.P.Y
		z.V.Z = (b.P.Z + b.V.Z) - z.P.Z
	} else {
		z.V.X = (b.V.X*mub + b.P.X) - z.P.X
		z.V.Y = (b.V.Y*mub + b.P.Y) - z.P.Y
		z.V.Z = (b.V.Z*mub + b.P.Z) - z.P.Z
	}
	return 1
}

// Intersection3DPlaneLine sets z to the intersecion of plane a and line b,
// then returns 1.
func Intersection3DPlaneLine(a *Plane, b *Line3D, z *Vector3D) int {
	// http://paulbourke.net/geometry/planeline/
	u := (a.A*b.P.X + a.B*b.P.Y + a.C*b.P.Z + a.D) / (a.A*b.V.X + a.B*b.V.Y + a.C*b.V.Z)
	z.X = b.P.X - u*b.V.X
	z.Y = b.P.Y - u*b.V.Y
	z.Z = b.P.Z - u*b.V.Z
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
	z.P.X = c1*a.A + c2*b.A
	z.P.Y = c1*a.B + c2*b.B
	z.P.Z = c1*a.C + c2*b.C
	z.V.X = a.B*b.C - a.C*b.B
	z.V.Y = a.C*b.A - a.A*b.C
	z.V.Z = a.A*b.B - a.B*b.A
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
	dot2 := a.A*b.V.X + a.B*b.V.Y + a.C*b.V.Z
	dot1 := a.A*b.P.X + a.B*b.P.Y + a.C*b.P.Z
	if FuzzyEqual(dot2, 0) {
		if FuzzyEqual(dot1, 0) {
			return -1
		}
		return 0
	}
	u := (dot1 + a.D) / dot2
	z.X = b.P.X - u*b.V.X
	z.Y = b.P.Y - u*b.V.Y
	z.Z = b.P.Z - u*b.V.Z
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
		if s*a.D*b.D < 0 || !FuzzyEqual(a.B, s*b.B) || !FuzzyEqual(a.C, s*b.C) ||
			!FuzzyEqual(b.D*b.D*n1n1, a.D*a.D*n2n2) {
			return 0
		}
		return -1
	}
	n1n2 := a.A*b.A + a.B*b.B + a.C*b.C
	d := 1 / (n1n1*n2n2 - n1n2*n1n2)
	c1 := (b.D*n1n2 - a.D*n2n2) * d
	c2 := (a.D*n1n2 - b.D*n1n1) * d
	z.P.X = c1*a.A + c2*b.A
	z.P.Y = c1*a.B + c2*b.B
	z.P.Z = c1*a.C + c2*b.C
	z.V.X = cpx
	z.V.Y = cpy
	z.V.Z = cpz
	return 1
}

// Intersection3DFuzzyPlanePlanePlane sets z to the intersecion of 3 planes,
// then returns the number of intersections.
//
// Possible return values are:
// -3 if two planes are parallel and the third intersects at two lines.
// -2 if all three planes intersect at a line.
// -1 if all three planes are coincident.
// 0 if all planes are parallel (two could be coincident), z is untouched.
// 1 if the planes intersect at a point, z is set to the intersection point.
func Intersection3DFuzzyPlanePlanePlane(a, b, c *Plane, z *Vector3D) int {
	n1n1 := a.A*a.A + a.B*a.B + a.C*a.C
	n2n2 := b.A*b.A + b.B*b.B + b.C*b.C
	n3n3 := c.A*c.A + c.B*c.B + c.C*c.C

	// use cross products to check if plane normals are equal
	cpabx, cpaby, cpabz := a.B*b.C-a.C*b.B, a.C*b.A-a.A*b.C, a.A*b.B-a.B*b.A
	n1n2d := FuzzyEqual(cpabx*cpabx+cpaby*cpaby+cpabz*cpabz, 0)
	cpcax, cpcay, cpcaz := c.B*a.C-c.C*a.B, c.C*a.A-c.A*a.C, c.A*a.B-c.B*a.A
	n3n1d := FuzzyEqual(cpcax*cpcax+cpcay*cpcay+cpcaz*cpcaz, 0)

	// check if all planes are parallel
	if n1n2d && n3n1d {
		// check if all planes are coincident
		if FuzzyEqual(b.D*b.D*n1n1, a.D*a.D*n2n2) &&
			FuzzyEqual(c.D*c.D*n1n1, a.D*a.D*n3n3) {
			return -1
		} else {
			return 0
		}
	}

	cpbcx, cpbcy, cpbcz := b.B*c.C-b.C*c.B, b.C*c.A-b.A*c.C, b.A*c.B-b.B*c.A

	// check if lines from pair of plane intersections have the same direction
	if ldx, ldy, ldz := cpaby*cpbcz-cpabz*cpbcy, cpabz*cpbcx-cpabx*cpbcz,
		cpabx*cpbcy-cpaby*cpbcx; FuzzyEqual(ldx*ldx+ldy*ldy+ldz*ldz, 0) {
		// get point on each line
		n1n2 := a.A*b.A + a.B*b.B + a.C*b.C
		n2n3 := b.A*c.A + b.B*c.B + b.C*c.C
		d1 := 1 / (n1n1*n2n2 - n1n2*n1n2)
		d2 := 1 / (n2n2*n3n3 - n2n3*n2n3)
		c1 := (b.D*n1n2 - a.D*n2n2) * d1
		c2 := (a.D*n1n2 - b.D*n1n1) * d1
		c3 := (c.D*n2n3 - b.D*n3n3) * d2
		c4 := (b.D*n2n3 - c.D*n2n2) * d2
		p1x, p1y, p1z := c1*a.A+c2*b.A, c1*a.B+c2*b.B, c1*a.C+c2*b.C
		p2x, p2y, p2z := c3*b.A+c4*c.A, c3*b.B+c4*c.B, c3*b.C+c4*c.C

		// check if points lie on the third plane
		if FuzzyEqual(p1x*c.A+p1y*c.B+p1z*c.C+c.D, 0) &&
			FuzzyEqual(p2x*a.A+p2y*a.B+p2z*a.C+a.D, 0) {
			return -2
		}
	}

	// check for a pair of parallel planes resulting in 2 lines, all 3 parallel
	// and 2 coincident have been caught already
	if n1n2d || n3n1d || FuzzyEqual(cpbcx*cpbcx+cpbcy*cpbcy+cpbcz*cpbcz, 0) {
		return -3
	}

	// having ruled out all degenerate cases calculate intersection
	d := -1 / (a.A*cpbcx + a.B*cpbcy + a.C*cpbcz)
	z.X = (a.D*cpbcx + b.D*cpcax + c.D*cpabx) * d
	z.Y = (a.D*cpbcy + b.D*cpcay + c.D*cpaby) * d
	z.Z = (a.D*cpbcz + b.D*cpcaz + c.D*cpabz) * d
	return 1
}
