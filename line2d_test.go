package geometry

import (
	"math"
	"testing"
)

func TestLine2DCore(t *testing.T) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	if v := l.ToVector(); !v.Equal(&Vector2D{3, 4}) {
		t.Error("Line2D.ToVector2D")
	}
	if l.Length() != 5 {
		t.Error("Line2D.Length")
	}
	if l.LengthSquared() != 25 {
		t.Error("Line2D.LengthSquared")
	}
}

func Benchmark_Line2D_Length(b *testing.B) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	for i := 0; i < b.N; i++ {
		l.Length()
	}
}

func Benchmark_Line2D_LengthSquared(b *testing.B) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	for i := 0; i < b.N; i++ {
		l.LengthSquared()
	}
}

func Benchmark_Line2D_ToVector2D(b *testing.B) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	for i := 0; i < b.N; i++ {
		l.ToVector()
	}
}

func TestMidpoint2D(t *testing.T) {
	l := Line2D{Point2D{0, 0}, Point2D{1, 1}}
	if m := l.Midpoint(); !m.Equal(&Point2D{0.5, 0.5}) {
		t.Error("Line2D.Midpoint")
	}
}

func Benchmark_Line2D_Midpoint(b *testing.B) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	for i := 0; i < b.N; i++ {
		l.Midpoint()
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

func Benchmark_Line2D_Normal(b *testing.B) {
	l := Line2D{Point2D{1, 1}, Point2D{4, 5}}
	for i := 0; i < b.N; i++ {
		l.Normal()
	}
}

func TestLine2DSegmentPointDistance(t *testing.T) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	if l.SegmentPointDistance(&Point2D{0, 0}) != math.Sqrt2 {
		t.Error("SegmentPointDistance")
	}
	if l.SegmentPointDistance(&Point2D{1.5, 0}) != 1 {
		t.Error("SegmentPointDistance")
	}
	if l.SegmentPointDistance(&Point2D{3, 0}) != math.Sqrt2 {
		t.Error("SegmentPointDistance")
	}
}

func Benchmark_Line2D_SegmentPointDistance_P1(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p1 := &Point2D{0, 0}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistance(p1)
	}
}

func Benchmark_Line2D_SegmentPointDistance_P2(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p2 := &Point2D{1.5, 0}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistance(p2)
	}
}

func Benchmark_Line2D_SegmentPointDistance_P3(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p3 := &Point2D{3, 0}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistance(p3)
	}
}

func TestLine2DPointDistance(t *testing.T) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	if l.PointDistance(&Point2D{0, 0}) != 1 {
		t.Error("LinePointDistance")
	}
	if l.PointDistance(&Point2D{1.5, 0}) != 1 {
		t.Error("LinePointDistance")
	}
	if l.PointDistance(&Point2D{3, 0}) != 1 {
		t.Error("LinePointDistance")
	}
}

func Benchmark_Line2D_PointDistance(b *testing.B) {
	l := &Line2D{Point2D{1, 1}, Point2D{2, 1}}
	p1 := &Point2D{0, 0}
	for i := 0; i < b.N; i++ {
		l.PointDistance(p1)
	}
}

func TestLine2DSegmentIntersection(t *testing.T) {
	l1 := Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := Line2D{Point2D{0, 1}, Point2D{1, 0}}
	p, seg := l1.SegmentIntersection(&l2)
	if !p.Equal(&Point2D{0.5, 0.5}) || !seg {
		t.Error("Line2D.SegmentIntersection")
	}
}

func Benchmark_Line2D_SegmentIntersection_1(b *testing.B) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{0, 1}, Point2D{1, 0}}
	for i := 0; i < b.N; i++ {
		l1.SegmentIntersection(l2)
	}
}

func Benchmark_Line2D_SegmentIntersection_2(b *testing.B) {
	l1 := &Line2D{Point2D{0.5, 0.5}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{0, 0}, Point2D{1, 0}}
	for i := 0; i < b.N; i++ {
		l1.SegmentIntersection(l2)
	}
}

func TestLine2DIntersection(t *testing.T) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{0, 1}, Point2D{1, 0}}
	p := l1.Intersection(l2)
	if !p.Equal(&Point2D{0.5, 0.5}) {
		t.Error("Line2D.Intersection")
	}
}

func Benchmark_Line2D_Intersection(b *testing.B) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{0, 1}, Point2D{1, 0}}
	for i := 0; i < b.N; i++ {
		l1.Intersection(l2)
	}
}

func TestLine2DEqual(t *testing.T) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{1, 1}, Point2D{0, 0}}
	if !l1.Equal(l2) || !l1.Equal(l1) {
		t.Error("Line2D.Equal")
	}
}

func Benchmark_Line2D_Equal(b *testing.B) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{0, 1}, Point2D{1, 0}}
	for i := 0; i < b.N; i++ {
		l1.Equal(l2)
	}
}

func TestLine2DFuzzyEqual(t *testing.T) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{1, 1}, Point2D{0, 0}}
	l1.P1.X += 0.0000000000001
	if !l1.FuzzyEqual(l2) || !l1.FuzzyEqual(l1) {
		t.Error("Line2D.FuzzyEqual")
	}
	l2.P1.X += 0.000000000001
	if l1.FuzzyEqual(l2) || !l1.FuzzyEqual(l1) {
		t.Error("Line2D.FuzzyEqual")
	}
}

func Benchmark_Line2D_FuzzyEqual(b *testing.B) {
	l1 := &Line2D{Point2D{0, 0}, Point2D{1, 1}}
	l2 := &Line2D{Point2D{0, 1}, Point2D{1, 0}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
	}
}
