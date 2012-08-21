package geometry

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
