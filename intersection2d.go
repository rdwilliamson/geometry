package geometry

// Intersection2DFuzzyLineLine sets point z to the intersection of a and b then
// returns the number of intersections.
//
// Possible return values are:
// -1 if the lines are coincident, z is untouched.
// 0 if the lines are parallel, z is untouched.
// 1 otherwise, z is set to the intersection of the two lines.
func Intersection2DFuzzyLineLine(a, b *Line2D, z *Vector2D) int {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
	adx, ady := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	bdx, bdy := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y
	d := bdy*adx - bdx*ady
	if FuzzyEqual(d, 0) {
		am, bm := ady/adx, bdy/bdx
		if FuzzyEqual(a.P1.Y-am*a.P1.X, b.P1.Y-bm*b.P1.X) {
			return -1
		}
		return 0
	}
	ua := (bdx*(a.P1.Y-b.P1.Y) - bdy*(a.P1.X-b.P1.X)) / d
	z.X = a.P1.X + ua*adx
	z.Y = a.P1.Y + ua*ady
	return 1
}

// Intersection2DLineLine sets point z to the intersection of a and b then
// returns z.
func Intersection2DLineLine(a, b *Line2D, z *Vector2D) *Vector2D {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
	adx, ady := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	bdx, bdy := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y
	ua := (bdx*(a.P1.Y-b.P1.Y) - bdy*(a.P1.X-b.P1.X)) / (bdy*adx - bdx*ady)
	z.X = a.P1.X + ua*adx
	z.Y = a.P1.Y + ua*ady
	return z
}

// Intersection2DLineSegmentLineSegment determines the intersection of two line
// segments then returns the number of intersections.
//
// Possible return values are:
// -1 if part of the line segments are coincident, z is untouched.
// 0 if the intersection does not occure on both line segments, z is untouched.
// 1 if the intersection occures on both line segments, z is set to the
// intersection.
func Intersection2DLineSegmentLineSegment(a, b *Line2D, z *Vector2D) int {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
	adx, ady := a.P2.X-a.P1.X, a.P2.Y-a.P1.Y
	bdx, bdy := b.P2.X-b.P1.X, b.P2.Y-b.P1.Y
	d := 1 / (bdy*adx - bdx*ady)
	dx, dy := a.P1.X-b.P1.X, a.P1.Y-b.P1.Y
	ua := (bdx*dy - bdy*dx) * d
	ub := (adx*dy - ady*dx) * d
	if 0 <= ua && ua <= 1 && 0 <= ub && ub <= 1 {
		z.X = a.P1.X + ua*adx
		z.Y = a.P1.Y + ua*ady
		return 1
	}
	return 0
}
