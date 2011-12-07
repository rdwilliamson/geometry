package geometry

// TODO want common functions which don't care how a line is defined
type Line2D interface {
	// Angle of the line in the range of [-pi/2 pi/2].
	Angle() float64
	// Distance a point is from the line.
	PointDistance(p Point2D)
	// Intersection point of two lines, may be at infinity.
	Intersection(l Line2D)
}
