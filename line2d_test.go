package geometry

import (
	"math"
	"testing"
)

func TestLine2DCore(t *testing.T) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	if v := l.ToVector2D(); !v.Equal(&Vector2D{3, 4}) {
		t.Error("Line2D.ToVector2D")
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
	if m := l.Midpoint(); !m.Equal(&Point2D{0.5, 0.5}) {
		t.Error("Line2D.Midpoint")
	}
}

func TestLine2DNormal(t *testing.T) {
	l := Line2D{Point2D{0, 0}, Point2D{2, 1}}
	n := l.Normal()
	la := math.Atan2(l.P2.Y-l.P1.Y, l.P2.X-l.P1.X)
	na := math.Atan2(n.Y, n.X)
	if l.Length() != n.Length() || math.Abs(la-na) != math.Pi/2 {
		t.Error("Line2D.Normal")
	}
}

func TestLine2DIntersection(t *testing.T) {
	l1 := Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := Line2D{Point2D{0, 1}, Point2D{1, 0}}
	p, _ := l1.Intersection(&l2)
	if !p.Equal(&Point2D{0.5, 0.5}) {
		t.Error("Line2D.Intersection")
	}
}

func TestLineSegmentPointDistance2D(t *testing.T) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	if LineSegmentPointDistance2D(l, &Point2D{0, 0}) != math.Sqrt2 {
		t.Error("LineSegmentPointDistance2D")
	}
	if LineSegmentPointDistance2D(l, &Point2D{1.5, 0}) != 1 {
		t.Error("LineSegmentPointDistance2D")
	}
	if LineSegmentPointDistance2D(l, &Point2D{3, 0}) != math.Sqrt2 {
		t.Error("LineSegmentPointDistance2D")
	}
}

func BenchmarkLineSegmentPointDistance2D_P1(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p1 := &Point2D{0, 0}
	for i := 0; i < b.N; i++ {
		LineSegmentPointDistance2D(l, p1)
	}
}

func BenchmarkLineSegmentPointDistance2D_P2(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p2 := &Point2D{1.5, 0}
	for i := 0; i < b.N; i++ {
		LineSegmentPointDistance2D(l, p2)
	}
}

func BenchmarkLineSegmentPointDistance2D_P3(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p3 := &Point2D{3, 0}
	for i := 0; i < b.N; i++ {
		LineSegmentPointDistance2D(l, p3)
	}
}

func TestLinePointDistance2D(t *testing.T) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	if LinePointDistance2D(l, &Point2D{0, 0}) != 1 {
		t.Error("LinePointDistance2D")
	}
	if LinePointDistance2D(l, &Point2D{1.5, 0}) != 1 {
		t.Error("LinePointDistance2D")
	}
	if LinePointDistance2D(l, &Point2D{3, 0}) != 1 {
		t.Error("LinePointDistance2D")
	}
}

func BenchmarkLinePointDistance2D_P1(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p1 := &Point2D{0, 0}
	for i := 0; i < b.N; i++ {
		LinePointDistance2D(l, p1)
	}
}

func BenchmarkLinePointDistance2D_P2(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p2 := &Point2D{1.5, 0}
	for i := 0; i < b.N; i++ {
		LinePointDistance2D(l, p2)
	}
}

func BenchmarkLinePointDistance2D_P3(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p3 := &Point2D{3, 0}
	for i := 0; i < b.N; i++ {
		LinePointDistance2D(l, p3)
	}
}
