package geometry

import "testing"

func TestPoint2DAdd(t *testing.T) {
	r, p1, p2 := &Point2D{0, 0}, &Point2D{1, 2}, &Point2D{3, 4}
	if !r.Add(p1, p2).Equal(&Point2D{4, 6}) {
		t.Error("Point2D.Add")
	}
}

func Benchmark_Point2D_Add(b *testing.B) {
	r, p1, p2 := &Point2D{0, 0}, &Point2D{1, 2}, &Point2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Add(p1, p2)
	}
}

func TestPoint2DDist(t *testing.T) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	if p1.Dist(p2) != 1 {
		t.Error("Point2D.Dist")
	}
}

func Benchmark_Point2D_Dist(b *testing.B) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.Dist(p2)
	}
}

func TestPoint2DEqual(t *testing.T) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	if p1.Equal(p2) {
		t.Error("Point2D.Equal")
	}
	p2.X = 0
	if !p1.Equal(p2) {
		t.Error("Point2D.Equal")
	}
}

func Benchmark_Point2D_Equal(b *testing.B) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func TestPoint2DFuzzyEqual(t *testing.T) {
	p1, p2 := &Point2D{0, 0}, &Point2D{0.000000000001, 0}
	if p1.FuzzyEqual(p2) {
		t.Error("Point2D.FuzzyEqual")
	}
	p2.X = 0.0000000000001
	if !p1.FuzzyEqual(p2) {
		t.Error("Point2D.FuzzyEqual")
	}
}

func Benchmark_Point2D_FuzzyEqual(b *testing.B) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.FuzzyEqual(p2)
	}
}

func TestPoint2DSet(t *testing.T) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	if !p1.Set(p2).Equal(p2) {
		t.Error("Point2D.Set")
	}
}

func Benchmark_Point2D_Set(b *testing.B) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.Set(p2)
	}
}

func TestPoint2DSqDist(t *testing.T) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	if p1.SqDist(p2) != 1 {
		t.Error("Point2D.SqDist")
	}
}

func Benchmark_Point2D_SqDist(b *testing.B) {
	p1, p2 := &Point2D{0, 0}, &Point2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.SqDist(p2)
	}
}

func TestPoint2DSub(t *testing.T) {
	r, p1, p2 := &Point2D{0, 0}, &Point2D{1, 2}, &Point2D{3, 4}
	if !r.Sub(p1, p2).Equal(&Point2D{-2, -2}) {
		t.Error("Point2D.Sub")
	}
}

func Benchmark_Point2D_Sub(b *testing.B) {
	r, p1, p2 := &Point2D{0, 0}, &Point2D{1, 2}, &Point2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Sub(p1, p2)
	}
}
