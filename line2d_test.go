package geometry

import (
	"math"
	"testing"
)

func TestLine2DCore(t *testing.T) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	if !l.ToVector2D().Equal(Vector2D{3, 4}) {
		t.Error("Line2D.ToVector2D")
	}
	if l.Dx() != 3 {
		t.Error("Line2D.Dx")
	}
	if l.Dy() != 4 {
		t.Error("Line2D.Dy")
	}
	if l.Length() != 5 {
		t.Error("Line2D.Length")
	}
	if l.LengthSquared() != 25 {
		t.Error("Line2D.LengthSquared")
	}
}

func TestMidpoint2D(t *testing.T) {
	l := Line2D{Point2D{0, 0}, Point2D{1, 1}}
	if !l.Midpoint().Equal(Point2D{0.5, 0.5}) {
		t.Error("Line2D.Midpoint")
	}
}

func TestAngle2D(t *testing.T) {
	l := Line2D{Point2D{0, 0}, Point2D{1, 1}}
	if l.Angle() != math.Pi/4 {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{-1, 1}
	if l.Angle() != 3*math.Pi/4 {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{-1, -1}
	if l.Angle() != -3*math.Pi/4 {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{1, -1}
	if l.Angle() != -math.Pi/4 {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{1, 0}
	if l.Angle() != 0 {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{0, 1}
	if l.Angle() != math.Pi/2 {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{-1, 0}
	if l.Angle() != math.Pi {
		t.Error("Line2D.Angle")
	}
	l.P2 = Point2D{0, -1}
	if l.Angle() != -math.Pi/2 {
		t.Error("Line2D.Angle")
	}
}

func TestLinePointDistance2D(t *testing.T) {
	l := Line2D{Point2D{1, 1}, Point2D{2, 1}}
	if LinePointDistance2D(l, Point2D{0, 0}, true) != math.Sqrt2 {
		t.Error("LinePointDistance2D")
	}
	if LinePointDistance2D(l, Point2D{0, 0}, false) != 1 {
		t.Error("LinePointDistance2D")
	}
	if LinePointDistance2D(l, Point2D{3, 0}, true) != math.Sqrt2 {
		t.Error("LinePointDistance2D")
	}
}

func TestLine2DNormal(t *testing.T) {
	l := Line2D{Point2D{0, 0}, Point2D{2, 1}}
	n := l.Normal()
	na := math.Atan2(n.Y, n.X)
	if l.Length() != n.Length() || math.Abs(l.Angle()-na) != math.Pi/2 {
		t.Error("Line2D.Normal")
	}
}

func TestLine2DIntersection(t *testing.T) {
	l1 := Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := Line2D{Point2D{0, 1}, Point2D{1, 0}}
	p, seg := l1.Intersection(l2)
	if !p.Equal(Point2D{0.5, 0.5}) || !seg {
		t.Fatal("Line2D.Intersection", p, seg)
	}
}
