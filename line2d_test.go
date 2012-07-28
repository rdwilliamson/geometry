package geometry

import (
	"testing"
)

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

func TestLine2DMidpoint(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	if !l.Midpoint(v).Equal(&Vector2D{2, 3}) {
		t.Error("Line2D.Midpoint", v)
	}
}

func Benchmark_Line2D_Midpoint(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		l.Midpoint(v)
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
