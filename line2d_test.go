package geometry

import (
	"testing"
)

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
