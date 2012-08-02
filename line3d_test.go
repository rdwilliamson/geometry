package geometry

import (
	"math"
	"testing"
)

func TestLine3DEqual(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	l2 := &Line3D{Vector3D{5, 6, 7}, Vector3D{0, 1, 2}}
	if !l1.Equal(l2) {
		t.Error("Line3D.Equal")
	}
}

func Benchmark_Line3D_Equal(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	l2 := &Line3D{Vector3D{5, 6, 7}, Vector3D{0, 1, 2}}
	for i := 0; i < b.N; i++ {
		l1.Equal(l2)
	}
}

func TestLine3DFuzzyEqual(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	l2 := &Line3D{Vector3D{5, 6, 7}, Vector3D{0, 1, 2.00000000001}}
	if l1.FuzzyEqual(l2) {
		t.Error("Line3D.FuzzyEqual")
	}
	l2.P2.Z = 2.000000000001
	if !l1.FuzzyEqual(l2) {
		t.Error("Line3D.FuzzyEqual")
	}
}

func Benchmark_Line3D_FuzzyEqual(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	l2 := &Line3D{Vector3D{5, 6, 7}, Vector3D{0, 1, 2}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
	}
}

func TestLine3DLineBetween(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{3, 3, 3}}
	l2 := &Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 3}}
	r := &Line3D{}
	lb := &Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}
	if r.LineBetween(l1, l2); !r.SegmentEqual(lb) {
		t.Error("Line3D.LineBetween")
	}
	l1 = &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 = &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	lb = &Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}
	if r.LineBetween(l1, l2); !r.SegmentFuzzyEqual(lb) {
		t.Error("Line3D.LineBetween")
	}
}

func Benchmark_Line3D_LineBetween(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		r.LineBetween(l1, l2)
	}
}

func TestLine3DLength(t *testing.T) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{6, 5, 4}}
	if l.Length() != math.Sqrt(35) {
		t.Error("Line3D.Length")
	}
}

func Benchmark_Line3D_Length(b *testing.B) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{6, 5, 4}}
	for i := 0; i < b.N; i++ {
		l.Length()
	}
}

func TestLine3DLengthSquared(t *testing.T) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{6, 5, 4}}
	if l.LengthSquared() != 35 {
		t.Error("Line3D.LengthSquared")
	}
}

func Benchmark_Line3D_LengthSquared(b *testing.B) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{6, 5, 4}}
	for i := 0; i < b.N; i++ {
		l.LengthSquared()
	}
}

func TestLine3DSegmentEqual(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{1, 1, 1}, Vector3D{1, 2, 3}}
	if !l1.SegmentEqual(l2) {
		t.Error("Line3D.SegmentEqual")
	}
}

func Benchmark_Line3D_SegmentEqual(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{1, 1, 1}, Vector3D{1, 2, 3}}
	for i := 0; i < b.N; i++ {
		l1.SegmentEqual(l2)
	}
}

func TestLine3DSegmentFuzzyEqual(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{1, 1, 1}, Vector3D{1, 2, 3}}
	l1.P1.X += 0.0000000000001
	if !l1.SegmentFuzzyEqual(l2) || !l1.SegmentFuzzyEqual(l1) {
		t.Error("Line3D.SegmentFuzzyEqual")
	}
	l2.P1.X += 0.000000000001
	if l1.SegmentFuzzyEqual(l2) || !l1.SegmentFuzzyEqual(l1) {
		t.Error("Line3D.SegmentFuzzyEqual")
	}
}

func Benchmark_Line3D_SegmentFuzzyEqual(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{1, 1, 1}, Vector3D{1, 2, 3}}
	for i := 0; i < b.N; i++ {
		l1.SegmentFuzzyEqual(l2)
	}
}

func TestLine3DPointDistance(t *testing.T) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if l.PointDistance(p) != 1 {
		t.Error("Line3D.PointDistance")
	}
}

func Benchmark_Line3D_PointDistance(b *testing.B) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		l.PointDistance(p)
	}
}

func TestLine3DPointSquaredDistance(t *testing.T) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if l.PointSquaredDistance(p) != 1 {
		t.Error("Line3D.PointSquaredDistance")
	}
}

func Benchmark_Line3D_PointSquaredDistance(b *testing.B) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		l.PointSquaredDistance(p)
	}
}
