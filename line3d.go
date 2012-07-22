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
