package geometry

type Point2D struct {
	X, Y float64
}

func (p1 Point2D) Plus(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Point2D) Add(p2 Point2D) {
	p1.X += p2.X
	p2.Y += p2.Y
}

func (p1 Point2D) Minus(p2 Point2D) Point2D {
	return Point2D{p1.X - p2.X, p1.Y - p2.Y}
}

func (p1 *Point2D) Subtract(p2 Point2D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
}

func (p Point2D) Scaled(s float64) Point2D {
	return Point2D{p.X * s, p.Y * s}
}

func (p *Point2D) Scale(s float64) {
	p.X *= s
	p.Y *= s
}
