package geometry

import (
	"math"
	"testing"
)

func (v *Vector2D) toPositiveInf() *Vector2D {
	if math.IsInf(v.X, -1) {
		v.X = math.Inf(1)
	}
	if math.IsInf(v.Y, -1) {
		v.Y = math.Inf(1)
	}
	return v
}

func (a *Vector2D) nanEqual(b *Vector2D) bool {
	if math.IsNaN(a.X) && math.IsNaN(a.Y) && math.IsNaN(b.X) && math.IsNaN(b.Y) {
		return true
	}
	return a.Equal(b)
}

type intersection2DFuzzyLineLineData struct {
	l1, l2 Line2D
	p      Vector2D
	n      int
}

var intersection2DFuzzyLineLineValues = []intersection2DFuzzyLineLineData{
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, -1}},
		Vector2D{0.5, 0.5}, 1},
	{Line2D{Vector2D{0, 0}, Vector2D{-1, -1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, -1}},
		Vector2D{0.5, 0.5}, 1},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{1, 0}, Vector2D{1, 1}},
		Vector2D{}, 0},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{-1, -1}, Vector2D{1, 1}},
		Vector2D{}, -1},
}

func testIntersection2DFuzzyLineLine(d intersection2DFuzzyLineLineData, t *testing.T) {
	var p Vector2D
	if n := Intersection2DFuzzyLineLine(&d.l1, &d.l2, &p); n != d.n ||
		(d.n == 1 && !p.Equal(&d.p)) {
		if d.n == 1 {
			t.Error("Intersection2D.FuzzyLineLine", d.l1, d.l2, "want", d.n,
				d.p, "got", n, p)
		} else {
			t.Error("Intersection2D.FuzzyLineLine", d.l1, d.l2, "want", d.n,
				"got", n)
		}
	}
}

func TestIntersection2DFuzzyLineLine(t *testing.T) {
	for _, v := range intersection2DFuzzyLineLineValues {
		testIntersection2DFuzzyLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection2DFuzzyLineLine(v, t)
		v.l1.P.X, v.l1.P.Y = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y
		v.l1.V.X, v.l1.V.Y = -v.l1.V.X, -v.l1.V.Y
		testIntersection2DFuzzyLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection2DFuzzyLineLine(v, t)
		v.l1.P.X, v.l1.P.Y = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y
		v.l1.V.X, v.l1.V.Y = -v.l1.V.X, -v.l1.V.Y
		testIntersection2DFuzzyLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection2DFuzzyLineLine(v, t)
	}
}

func Benchmark_Intersection2DFuzzy_LineLine_Coincident(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineLine_Parallel(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{1, 0}, Vector2D{2, 1}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineLine(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineLine_Intersect(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineLine(l1, l2, p)
	}
}

type intersection2DFuzzyLineSegmentLineSegmentData struct {
	l1, l2 Line2D
	p      Vector2D
	n      int
}

var intersection2DFuzzyLineSegmentLineSegmentValues = []intersection2DFuzzyLineSegmentLineSegmentData{
	// intersect
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, -1}},
		Vector2D{0.5, 0.5}, 1},
	// intersect if they were lines
	{Line2D{Vector2D{0, 0}, Vector2D{-1, -1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, -1}},
		Vector2D{}, 0},
	// parallel but not coincident
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{1, 0}, Vector2D{1, 1}},
		Vector2D{}, 0},
	// coincident
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{-0.5, -0.5}, Vector2D{1, 1}},
		Vector2D{}, -1},
	// coincident if lines
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{-2, -2}, Vector2D{1, 1}},
		Vector2D{}, 0},
	// endpoints match
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{2, 2}, Vector2D{-1, -1}},
		Vector2D{1, 1}, 1},
}

func testIntersection2DFuzzyLineSegmentLineSegment(d intersection2DFuzzyLineSegmentLineSegmentData, t *testing.T) {
	var p Vector2D
	if n := Intersection2DFuzzyLineSegmentLineSegment(&d.l1, &d.l2, &p); n != d.n ||
		(d.n == 1 && !p.Equal(&d.p)) {
		if d.n == 1 {
			t.Error("Intersection2D.FuzzyLineSegmentLineSegment", d.l1, d.l2,
				"want", d.n, d.p, "got", n, p)
		} else {
			t.Error("Intersection2D.FuzzyLineSegmentLineSegment", d.l1, d.l2,
				"want", d.n, "got", n)
		}
	}
}

func TestIntersection2DFuzzyLineSegmentLineSegment(t *testing.T) {
	for _, v := range intersection2DFuzzyLineSegmentLineSegmentValues {
		testIntersection2DFuzzyLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection2DFuzzyLineSegmentLineSegment(v, t)
		v.l1.P.X, v.l1.P.Y = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y
		v.l1.V.X, v.l1.V.Y = -v.l1.V.X, -v.l1.V.Y
		testIntersection2DFuzzyLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection2DFuzzyLineSegmentLineSegment(v, t)
		v.l1.P.X, v.l1.P.Y = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y
		v.l1.V.X, v.l1.V.Y = -v.l1.V.X, -v.l1.V.Y
		testIntersection2DFuzzyLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection2DFuzzyLineSegmentLineSegment(v, t)
	}
}

func Benchmark_Intersection2DFuzzy_LineSegmentLineSegment_Endpoint(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{-1, -1}, Vector2D{0, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineSegmentLineSegment(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineSegmentLineSegment_Coincident(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{-0.5, -0.5}, Vector2D{0.5, 0.5}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineSegmentLineSegment(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineSegmentLineSegment_Parallel(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{1, 0}, Vector2D{2, 1}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineSegmentLineSegment(l1, l2, p)
	}
}

func Benchmark_Intersection2DFuzzy_LineSegmentLineSegment_Intersect(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DFuzzyLineSegmentLineSegment(l1, l2, p)
	}
}

type intersection2DLineLineData struct {
	l1, l2 Line2D
	p      Vector2D
}

var intersection2DLineLineValues = []intersection2DLineLineData{
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, -1}},
		Vector2D{0.5, 0.5}},
	{Line2D{Vector2D{0, 0}, Vector2D{-1, -1}},
		Line2D{Vector2D{0, 1}, Vector2D{1, -1}},
		Vector2D{0.5, 0.5}},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{1, 0}, Vector2D{1, 1}},
		Vector2D{math.Inf(1), math.Inf(1)}},
	{Line2D{Vector2D{0, 0}, Vector2D{1, 1}},
		Line2D{Vector2D{-2, -2}, Vector2D{1, 1}},
		Vector2D{math.NaN(), math.NaN()}},
}

func testIntersection2DLineLine(d intersection2DLineLineData, t *testing.T) {
	var p Vector2D
	if Intersection2DLineLine(&d.l1, &d.l2, &p); !p.toPositiveInf().nanEqual(&d.p) {
		t.Error("Intersection2D.LineLine", d.l1, d.l2, "want", d.p, "got", p)
	}
	if Intersection2DLineLine(&d.l2, &d.l1, &p); !p.toPositiveInf().nanEqual(&d.p) {
		t.Error("Intersection2D.LineLine", d.l2, d.l1, "want", d.p, "got", p)
	}
}

func TestIntersection2DLineLine(t *testing.T) {
	for _, v := range intersection2DLineLineValues {
		testIntersection2DLineLine(v, t)
		v.l1.P.X, v.l1.P.Y = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y
		v.l1.V.X, v.l1.V.Y = -v.l1.V.X, -v.l1.V.Y
		testIntersection2DLineLine(v, t)
		v.l2.P.X, v.l2.P.Y = v.l2.P.X+v.l2.V.X, v.l2.P.Y+v.l2.V.Y
		v.l2.V.X, v.l2.V.Y = -v.l2.V.X, -v.l2.V.Y
		testIntersection2DLineLine(v, t)
	}
}

func Benchmark_Intersection2D_LineLine(b *testing.B) {
	l1 := &Line2D{Vector2D{0, 0}, Vector2D{1, 1}}
	l2 := &Line2D{Vector2D{0, 1}, Vector2D{1, 0}}
	p := &Vector2D{}
	for i := 0; i < b.N; i++ {
		Intersection2DLineLine(l1, l2, p)
	}
}
