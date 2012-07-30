package geometry

import (
	"math"
	"testing"
)

func TestNewLine2D(t *testing.T) {
	if !NewLine2D(1, 2, 3, 4).SegmentEqual(&Line2D{Vector2D{1, 2}, Vector2D{3, 4}}) {
		t.Error("NewLine2D")
	}
}

func Benchmark_Line2D_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewLine2D(1, 2, 3, 4)
	}
}

func TestLine2DAngleDistance(t *testing.T) {
	l := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	p := &Vector2D{1, 0.5}
	if !FuzzyEqual(l.AngleDistance(p), math.Pi/4) {
		t.Error("Line2D.AngleDistance")
	}
	p = &Vector2D{0, 0.5}
	if !FuzzyEqual(l.AngleDistance(p), math.Pi/4) {
		t.Error("Line2D.AngleDistance")
	}
	l = &Line2D{Vector2D{1, 1}, Vector2D{0, 0}}
	if !FuzzyEqual(l.AngleDistance(p), math.Pi/4) {
		t.Error("Line2D.AngleDistance")
	}
	p = &Vector2D{1, 0.5}
	if !FuzzyEqual(l.AngleDistance(p), math.Pi/4) {
		t.Error("Line2D.AngleDistance")
	}
}

func Benchmark_Line2D_AngleDistance(b *testing.B) {
	l := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	p := &Vector2D{1, 0.5}
	for i := 0; i < b.N; i++ {
		l.AngleDistance(p)
	}
}

func TestLine2DAngleCosSquaredDistance(t *testing.T) {
	l := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	p := &Vector2D{1, 0.5}
	if l.AngleCosSquaredDistance(p) != 0.5 {
		t.Error("Line2D.AngleCosSquaredDistance")
	}
	p = &Vector2D{0, 0.5}
	if l.AngleCosSquaredDistance(p) != 0.5 {
		t.Error("Line2D.AngleCosSquaredDistance")
	}
	l = &Line2D{Vector2D{1, 1}, Vector2D{0, 0}}
	if l.AngleCosSquaredDistance(p) != 0.5 {
		t.Error("Line2D.AngleCosSquaredDistance")
	}
	p = &Vector2D{1, 0.5}
	if l.AngleCosSquaredDistance(p) != 0.5 {
		t.Error("Line2D.AngleCosSquaredDistance")
	}
}

func Benchmark_Line2D_AngleCosSquaredDistance(b *testing.B) {
	l := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	p := &Vector2D{1, 0.5}
	for i := 0; i < b.N; i++ {
		l.AngleCosSquaredDistance(p)
	}
}

func TestLine2DEqual(t *testing.T) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{-3, -2}, Vector2D{5, 6}}
	if !l1.Equal(l2) {
		t.Error("Line2D.Equal")
	}
}

func Benchmark_Line2D_Equal(b *testing.B) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{-3, -2}, Vector2D{5, 6}}
	for i := 0; i < b.N; i++ {
		l1.Equal(l2)
	}
}

func TestLine2DFuzzyEqual(t *testing.T) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{-3, -2}, Vector2D{5, 6.00000000001}}
	if l1.FuzzyEqual(l2) {
		t.Error("Line2D.FuzzyEqual")
	}
	l2.P2.Y = 6.000000000001
	if !l1.FuzzyEqual(l2) {
		t.Error("Line2D.FuzzyEqual")
	}
}

func Benchmark_Line2D_FuzzyEqual(b *testing.B) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{-3, -2}, Vector2D{5, 6}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
	}
}

func TestLine2DIntersection(t *testing.T) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	if !l1.Intersection(l2, p).Equal(&Vector2D{0.5, 0.5}) {
		t.Error("Line2D.Intersection")
	}
}

func Benchmark_Line2D_Intersection(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		l1.Intersection(l2, p)
	}
}

func TestLine2DLength(t *testing.T) {
	l := &Line2D{Vector2D{1, 2}, Vector2D{4, 6}}
	if l.Length() != 5 {
		t.Error("Line2D.Length")
	}
}

func Benchmark_Line2D_Length(b *testing.B) {
	l := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l.Length()
	}
}

func TestLine2DLengthSquared(t *testing.T) {
	l := &Line2D{Vector2D{1, 2}, Vector2D{4, 6}}
	if l.LengthSquared() != 25 {
		t.Error("Line2D.LengthSquared")
	}
}

func Benchmark_Line2D_LengthSquared(b *testing.B) {
	l := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l.LengthSquared()
	}
}

func TestLine2DMidpoint(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	if !l.Midpoint(v).Equal(&Vector2D{2, 3}) {
		t.Error("Line2D.Midpoint")
	}
}

func Benchmark_Line2D_Midpoint(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		l.Midpoint(v)
	}
}

func TestLine2DNormal(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 1}, Vector2D{3, 1}}, &Vector2D{}
	if !l.Normal(v).Equal(&Vector2D{0, -2}) {
		t.Error("Line2D.Normal", v)
	}
}

func Benchmark_Line2D_Normal(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 1}, Vector2D{3, 1}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		l.Normal(v)
	}
}

func TestLine2DPointDistance(t *testing.T) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	if l.PointDistance(p) != 1 {
		t.Error("Line2D.PointDistance")
	}
}

func Benchmark_Line2D_PointDistance(b *testing.B) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	for i := 0; i < b.N; i++ {
		l.PointDistance(p)
	}
}

func TestLine2DPointDistanceSquared(t *testing.T) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	if l.PointDistanceSquared(p) != 1 {
		t.Error("Line2D.PointDistanceSquared")
	}
}

func Benchmark_Line2D_PointDistanceSquared(b *testing.B) {
	l, p := &Line2D{Vector2D{0, 0}, Vector2D{1, 0}}, &Vector2D{0, 1}
	for i := 0; i < b.N; i++ {
		l.PointDistanceSquared(p)
	}
}

func TestLine2DSegmentEqual(t *testing.T) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	if !l1.SegmentEqual(l2) {
		t.Error("Line2D.SegmentEqual")
	}
}

func Benchmark_Line2D_SegmentEqual(b *testing.B) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l1.SegmentEqual(l2)
	}
}

func TestLine2DSegmentFuzzyEqual(t *testing.T) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4.000000000001}}
	if l1.SegmentFuzzyEqual(l2) {
		t.Error("Line2D.SegmentFuzzyEqual")
	}
	l2.P2.Y = 4.0000000000001
	if !l1.SegmentFuzzyEqual(l2) {
		t.Error("Line2D.SegmentFuzzyEqual")
	}
}

func Benchmark_Line2D_SegmentFuzzyEqual(b *testing.B) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{2, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l1.SegmentFuzzyEqual(l2)
	}
}

func TestLine2DSegmentIntersection(t *testing.T) {
	l1 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	if l1.SegmentIntersection(l2, p) || !p.Equal(&Vector2D{0.5, 0.5}) {
		t.Error("Line2D.SegmentIntersection")
	}
	l1.P1 = Vector2D{1, 1}
	if !l1.SegmentIntersection(l2, p) || !p.Equal(&Vector2D{0.5, 0.5}) {
		t.Error("Line2D.SegmentIntersection")
	}
}

func Benchmark_Line2D_SegmentIntersection(b *testing.B) {
	l1 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		l1.SegmentIntersection(l2, p)
	}
}

func TestLine2DSegmentPointDistance(t *testing.T) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{-1, 0}
	if l.SegmentPointDistance(p) != math.Sqrt2 {
		t.Error("Line2D.SegmentPointDistance")
	}
	p.X = 0.5
	if l.SegmentPointDistance(p) != 1 {
		t.Error("Line2D.SegmentPointDistance")
	}
	p.X = 2
	if l.SegmentPointDistance(p) != math.Sqrt2 {
		t.Error("Line2D.SegmentPointDistance")
	}
}

func Benchmark_Line2D_SegmentPointDistance(b *testing.B) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{0.5, 0}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistance(p)
	}
}

func TestLine2DSegmentPointDistanceSquared(t *testing.T) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{-1, 0}
	if l.SegmentPointDistanceSquared(p) != 2 {
		t.Error("Line2D.SegmentPointDistanceSquared")
	}
	p.X = 0.5
	if l.SegmentPointDistanceSquared(p) != 1 {
		t.Error("Line2D.SegmentPointDistanceSquared")
	}
	p.X = 2
	if l.SegmentPointDistanceSquared(p) != 2 {
		t.Error("Line2D.SegmentPointDistanceSquared")
	}
}

func Benchmark_Line2D_SegmentPointDistanceSquared(b *testing.B) {
	l := &Line2D{Vector2D{0, 1}, Vector2D{1, 1}}
	p := &Vector2D{0.5, 0}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistanceSquared(p)
	}
}

func TestLine2DSet(t *testing.T) {
	l1, l2 := &Line2D{}, &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	if !l1.Set(l2).SegmentEqual(l2) {
		t.Error("Line2D.Set")
	}
}

func Benchmark_Line2D_Set(b *testing.B) {
	l1, l2 := &Line2D{}, &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l1.Set(l2)
	}
}

func TestLine2DToVector(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	if !l.ToVector(v).Equal(&Vector2D{2, 2}) {
		t.Error("Line2D.ToVector")
	}
}

func Benchmark_Line2D_ToVector(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		l.ToVector(v)
	}
}
