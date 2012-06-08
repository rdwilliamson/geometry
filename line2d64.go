package geometry

type Line2D64 struct {
	P1, P2 Point2D64
}

func (l Line2D64) ToVector2D64() Vector2D64 {
	return Vector2D64{l.P2.X - l.P1.X, l.P2.Y - l.P1.X}
}
