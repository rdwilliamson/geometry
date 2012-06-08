package geometry

type Point2D64 struct {
	X, Y float64
}

// func (p *Point2D64) ToPoint2D() Point2D {
// 	return Point2D{int(p.X), int(p.Y)}
// }

func (p *Point2D64) ToPoint2D32() Point2D32 {
	return Point2D32{float32(p.X), float32(p.Y)}
}

func (p1 Point2D64) Plus(p2 Point2D64) Point2D64 {
	return Point2D64{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Point2D64) Add(p2 Point2D64) {
	p1.X += p2.X
	p2.Y += p2.Y
}

func (p1 Point2D64) Minus(p2 Point2D64) Point2D64 {
	return Point2D64{p1.X - p2.X, p1.Y - p2.Y}
}

func (p1 *Point2D64) Subtract(p2 Point2D64) {
	p1.X -= p2.X
	p1.Y -= p2.Y
}

func (p Point2D64) Scaled(s float64) Point2D64 {
	return Point2D64{p.X * s, p.Y * s}
}

func (p *Point2D64) Scale(s float64) {
	p.X *= s
	p.Y *= s
}
