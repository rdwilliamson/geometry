package geometry

// Triangle2D represents a triangle by its three vertices.
type Triangle2D struct {
	P1, P2, P3 Point2D
}

// Orthocenter returns the orthocenter of the triangle.
// From http://2000clicks.com/MathHelp/GeometryTriangleCenterOrthocenter.aspx
func (t Triangle2D) Orthocenter() Point2D {
	a := t.P1.X
	b := t.P1.Y
	c := t.P2.X
	d := t.P2.Y
	e := t.P3.X
	f := t.P3.Y

	num1 := (d-f)*b*b + (f-b)*d*d + (b-d)*f*f + a*b*(c-e) + c*d*(e-a) + e*f*(a-c)
	num2 := (e-c)*a*a + (a-e)*c*c + (c-a)*e*e + a*b*(f-d) + c*d*(b-f) + e*f*(d-b)
	denom := 1.0 / (b*c + d*e + f*a - c*f - b*e - a*d)
	return Point2D{num1*denom, num2*denom}
}
