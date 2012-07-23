package geometry

import "testing"

func TestNewVector2D(t *testing.T) {
	if !NewVector2D(1, 2).Equal(&Vector2D{1, 2}) {
		t.Error("NewVector2D")
	}
}

func Benchmark_Vector2D_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewVector2D(0, 0)
	}
}

func TestVector2DAdd(t *testing.T) {
	r, p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 2}, &Vector2D{3, 4}
	if !r.Add(p1, p2).Equal(&Vector2D{4, 6}) {
		t.Error("Vector2D.Add")
	}
}

func Benchmark_Vector2D_Add(b *testing.B) {
	r, p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Add(p1, p2)
	}
}

func TestVector2DDistance(t *testing.T) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	if p1.Distance(p2) != 1 {
		t.Error("Vector2D.Distance")
	}
}

func Benchmark_Vector2D_Distance(b *testing.B) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.Distance(p2)
	}
}

func TestVector2DDistanceSquared(t *testing.T) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	if p1.DistanceSquared(p2) != 1 {
		t.Error("Vector2D.DistanceSquared")
	}
}

func Benchmark_Vector2D_DistanceSquared(b *testing.B) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.DistanceSquared(p2)
	}
}

func TestVector2DEqual(t *testing.T) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	if p1.Equal(p2) {
		t.Error("Vector2D.Equal")
	}
	p2.X = 0
	if !p1.Equal(p2) {
		t.Error("Vector2D.Equal")
	}
}

func Benchmark_Vector2D_Equal(b *testing.B) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func TestVector2DFuzzyEqual(t *testing.T) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{0.000000000001, 0}
	if p1.FuzzyEqual(p2) {
		t.Error("Vector2D.FuzzyEqual")
	}
	p2.X = 0.0000000000001
	if !p1.FuzzyEqual(p2) {
		t.Error("Vector2D.FuzzyEqual")
	}
}

func Benchmark_Vector2D_FuzzyEqual(b *testing.B) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.FuzzyEqual(p2)
	}
}

func TestVector2DSet(t *testing.T) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	if !p1.Set(p2).Equal(p2) {
		t.Error("Vector2D.Set")
	}
}

func Benchmark_Vector2D_Set(b *testing.B) {
	p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 0}
	for i := 0; i < b.N; i++ {
		p1.Set(p2)
	}
}

func TestVector2DSubtract(t *testing.T) {
	r, p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 2}, &Vector2D{3, 4}
	if !r.Subtract(p1, p2).Equal(&Vector2D{-2, -2}) {
		t.Error("Vector2D.Subtract")
	}
}

func Benchmark_Vector2D_Subtract(b *testing.B) {
	r, p1, p2 := &Vector2D{0, 0}, &Vector2D{1, 2}, &Vector2D{3, 4}
	for i := 0; i < b.N; i++ {
		r.Subtract(p1, p2)
	}
}
