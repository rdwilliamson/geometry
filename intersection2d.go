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
	d := b.V.Y*a.V.X - b.V.X*a.V.Y
	if FuzzyEqual(d, 0) {
		am, bm := a.V.Y/a.V.X, b.V.Y/b.V.X
		if FuzzyEqual(a.P.Y-am*a.P.X, b.P.Y-bm*b.P.X) {
			return -1
		}
		return 0
	}
	ua := (b.V.X*(a.P.Y-b.P.Y) - b.V.Y*(a.P.X-b.P.X)) / d
	z.X = a.P.X + ua*a.V.X
	z.Y = a.P.Y + ua*a.V.Y
	return 1
}

// Intersection2DFuzzyLineSegmentLineSegment determines the intersection of two
// line segments then returns the number of intersections.
//
// Possible return values are:
// -1 if part of the line segments are coincident, z is untouched.
// 0 if the intersection does not occure on both line segments, z is untouched.
// 1 if the intersection occures on both line segments, z is set to the
// intersection.
func Intersection2DFuzzyLineSegmentLineSegment(a, b *Line2D, z *Vector2D) int {
	d := (b.V.Y*a.V.X - b.V.X*a.V.Y)
	if FuzzyEqual(d, 0) {
		// slopes are the same, parallel or coincident
		am, bm := a.V.Y/a.V.X, b.V.Y/b.V.X
		if !FuzzyEqual(a.P.Y-am*a.P.X, b.P.Y-bm*b.P.X) {
			// parallel
			return 0
		}
		// check if endpoints are equal
		bp2x, bp2y := b.P.X+b.V.X, b.P.Y+b.V.Y
		if (FuzzyEqual(a.P.X, b.P.X) && FuzzyEqual(a.P.Y, b.P.Y)) ||
			(FuzzyEqual(a.P.X, bp2x) && FuzzyEqual(a.P.Y, bp2y)) {
			z.X = a.P.X
			z.Y = a.P.Y
			return 1
		}
		ap2x, ap2y := a.P.X+a.V.X, a.P.Y+a.V.Y
		if (FuzzyEqual(ap2x, b.P.X) && FuzzyEqual(ap2y, b.P.Y)) ||
			(FuzzyEqual(ap2x, bp2x) && FuzzyEqual(ap2y, bp2y)) {
			z.X = ap2x
			z.Y = ap2y
			return 1
		}
		// check for overlap
		var x1, x2 float64
		if a.P.X < ap2x {
			x1, x2 = a.P.X, ap2x
		} else {
			x1, x2 = ap2x, a.P.X
		}
		if (x1 < b.P.X && b.P.X < x2) || (x1 < bp2x && bp2x < x2) {
			return -1
		}
		// coincident if they were lines
		return 0
	}
	dx, dy := a.P.X-b.P.X, a.P.Y-b.P.Y
	d = 1 / d
	ua := (b.V.X*dy - b.V.Y*dx) * d
	ub := (a.V.X*dy - a.V.Y*dx) * d
	if 0 <= ua && ua <= 1 && 0 <= ub && ub <= 1 {
		z.X = a.P.X + ua*a.V.X
		z.Y = a.P.Y + ua*a.V.Y
		return 1
	}
	return 0
}

// Intersection2DLineLine sets point z to the intersection of a and b and
// returns 1.
func Intersection2DLineLine(a, b *Line2D, z *Vector2D) int {
	// http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
	ua := (b.V.X*(a.P.Y-b.P.Y) - b.V.Y*(a.P.X-b.P.X)) / (b.V.Y*a.V.X - b.V.X*a.V.Y)
	z.X = a.P.X + ua*a.V.X
	z.Y = a.P.Y + ua*a.V.Y
	return 1
}
