package geometry

// A Plane represents a plane by ax + by + cz + d = 0. A, B, and C represent
// the normal direction only (could have any magnitude). D divided by the
// magnitude of the normal represents the distance of the plane from the
// origin and the sign, if positive, signifies the origin is in the half space
// determined by the direction of the normal or the other half space if
// negative.
type Plane struct {
	A, B, C, D float64
}
