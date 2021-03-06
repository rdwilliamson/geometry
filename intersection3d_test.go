package geometry

import (
	"testing"
)

type intersection3DLineLineData struct {
	l1, l2, lb Line3D
}

var intersection3DLineLineValues = []intersection3DLineLineData{
	{Line3D{Vector3D{1, 2, 1}, Vector3D{2, 1, 2}}, Line3D{Vector3D{1, 2, 1}, Vector3D{0, 0, 2}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{0, 0, 0}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{4, 2, 4}}, Line3D{Vector3D{5, 6, 1}, Vector3D{-4, -2, 4}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{-0.8, 1.6, 0}}},
	{Line3D{Vector3D{5, 4, 5}, Vector3D{-2, -1, -2}}, Line3D{Vector3D{1, 2, 5}, Vector3D{0, 0, -2}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{0, 0, 0}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{-4, -2, -4}}, Line3D{Vector3D{5, 6, 1}, Vector3D{4, 2, -4}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{-0.8, 1.6, 0}}},
}

func testIntersection3DLineLine(d intersection3DLineLineData, t *testing.T) {
	var l Line3D
	if Intersection3DLineLine(&d.l1, &d.l2, &l); !l.SegmentFuzzyEqual(&d.lb) {
		t.Error("Intersection3D.LineLine", d.l1, d.l2, "want", d.lb, "got", l)
	}
}

func TestIntersection3DLineLine(t *testing.T) {
	for _, v := range intersection3DLineLineValues {
		testIntersection3DLineLine(v, t)
		v.l1.P.X, v.l1.P.Y, v.l1.P.Z = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y, v.l1.P.Z+v.l1.V.Z
		v.l1.V.X, v.l1.V.Y, v.l1.V.Z = -v.l1.V.X, -v.l1.V.Y, -v.l1.V.Z
		testIntersection3DLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineLine(v, t)
		v.l1.P.X, v.l1.P.Y, v.l1.P.Z = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y, v.l1.P.Z+v.l1.V.Z
		v.l1.V.X, v.l1.V.Y, v.l1.V.Z = -v.l1.V.X, -v.l1.V.Y, -v.l1.V.Z
		testIntersection3DLineLine(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineLine(v, t)
	}
}

func Benchmark_Intersection3D_LineLine(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DLineLine(l1, l2, r)
	}
}

type intersection3DLineSphereData struct {
	l      Line3D
	s      Sphere
	i1, i2 Vector3D
	n      int
}

var intersection3DLineSphereValues = []intersection3DLineSphereData{
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Sphere{Vector3D{}, 1}, Vector3D{1, 0, 0}, Vector3D{-1, 0, 0}, 2},
	{Line3D{Vector3D{0, 1, 0}, Vector3D{1, 0, 0}}, Sphere{Vector3D{}, 1}, Vector3D{0, 1, 0}, Vector3D{0, 0, 0}, 1},
	{Line3D{Vector3D{0, 2, 0}, Vector3D{1, 0, 0}}, Sphere{Vector3D{}, 1}, Vector3D{0, 0, 0}, Vector3D{0, 0, 0}, 0},
}

func testIntersection3DLineSphere(d intersection3DLineSphereData, t *testing.T) {
	var i1, i2 Vector3D
	n := Intersection3DLineSphere(&d.l, &d.s, &i1, &i2)
	if n != d.n {
		t.Error("Intersection3D.LineSphere", d.l, d.s, "want", d.n, "got", n)
		return
	}
	if n == 0 {
		return
	}
	if n == 1 {
		if d.i1.FuzzyEqual(&i1) {
			return
		}
		t.Error("Intersection3D.LineSphere", d.l, d.s, "want", d.i1, "got", i1)
		return
	}
	if d.i1.FuzzyEqual(&i1) && d.i2.FuzzyEqual(&i2) {
		return
	}
	if d.i1.FuzzyEqual(&i2) && d.i2.FuzzyEqual(&i1) {
		return
	}
	t.Error("Intersection3D.LineSphere", d.l, d.s, "want", d.i1, d.i2, "got", i1, i2)
}

func TestIntersction3DLineSphere(t *testing.T) {
	for _, v := range intersection3DLineSphereValues {
		testIntersection3DLineSphere(v, t)
	}
}

func Benchmark_Intersection3D_LineSphere(b *testing.B) {
	l := Line3D{Vector3D{}, Vector3D{1, 0, 0}}
	s := Sphere{Vector3D{}, 1}
	var i1, i2 Vector3D
	for i := 0; i < b.N; i++ {
		Intersection3DLineSphere(&l, &s, &i1, &i2)
	}
}

type intersection3DLineSegmentLineSegmentData struct {
	l1, l2, lb Line3D
}

var intersection3DLineSegmentLineSegmentValues = []intersection3DLineSegmentLineSegmentData{
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{4, 2, 4}}, Line3D{Vector3D{1, 2, -1}, Vector3D{0, 0, 4}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{0, 0, 0}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{-2, -1, -2}}, Line3D{Vector3D{1, 2, -1}, Vector3D{0, 0, 4}},
		Line3D{Vector3D{-1, 1, -1}, Vector3D{2, 1, 2}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{4, 2, 4}}, Line3D{Vector3D{1, 2, -1}, Vector3D{0, 0, -2}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{0, 0, -2}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{-2, -1, -2}}, Line3D{Vector3D{1, 2, -1}, Vector3D{0, 0, -2}},
		Line3D{Vector3D{-1, 1, -1}, Vector3D{2, 1, 0}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{4, 2, 4}}, Line3D{Vector3D{5, 6, 1}, Vector3D{-4, -2, 4}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{-0.8, 1.6, 0}}},
}

func testIntersection3DLineSegmentLineSegment(d intersection3DLineSegmentLineSegmentData, t *testing.T) {
	var l Line3D
	if Intersection3DLineSegmentLineSegment(&d.l1, &d.l2, &l); !l.SegmentFuzzyEqual(&d.lb) {
		t.Error("Intersection3D.LineSegmentLineSegment", d.l1, d.l2, "want",
			d.lb, "got", l)
	}
}

func TestIntersection3DLineSegmentLineSegment(t *testing.T) {
	for _, v := range intersection3DLineSegmentLineSegmentValues {
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1.P.X, v.l1.P.Y, v.l1.P.Z = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y, v.l1.P.Z+v.l1.V.Z
		v.l1.V.X, v.l1.V.Y, v.l1.V.Z = -v.l1.V.X, -v.l1.V.Y, -v.l1.V.Z
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1.P.X, v.l1.P.Y, v.l1.P.Z = v.l1.P.X+v.l1.V.X, v.l1.P.Y+v.l1.V.Y, v.l1.P.Z+v.l1.V.Z
		v.l1.V.X, v.l1.V.Y, v.l1.V.Z = -v.l1.V.X, -v.l1.V.Y, -v.l1.V.Z
		testIntersection3DLineSegmentLineSegment(v, t)
		v.l1, v.l2 = v.l2, v.l1
		testIntersection3DLineSegmentLineSegment(v, t)
	}
}

func Benchmark_Intersection3D_LineSegmentLineSegment(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DLineSegmentLineSegment(l1, l2, r)
	}
}

func TestIntersection3DPlaneLine(t *testing.T) {
	p := &Plane{1, -2, 1, -7}
	l := &Line3D{Vector3D{5, -1, -1}, Vector3D{3, -1, 1}}
	want, got := &Vector3D{11.0 / 2.0, -7.0 / 6.0, -5.0 / 6.0}, &Vector3D{}
	if Intersection3DPlaneLine(p, l, got); !got.Equal(want) {
		t.Error("Intersection3D.PlaneLine", *p, *l, "want", want, "got", got)
	}
}

func Benchmark_Intersection3D_PlaneLine(b *testing.B) {
	p := &Plane{1, -2, 1, -7}
	l := &Line3D{Vector3D{5, -1, -1}, Vector3D{8, -2, 0}}
	r := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DPlaneLine(p, l, r)
	}
}

func TestIntersection3DPlanePlane(t *testing.T) {
	p1, p2 := &Plane{1, 1, 1, 1}, &Plane{1, 2, 3, 4}
	got, want := &Line3D{}, &Line3D{Vector3D{3, -5, 1}, Vector3D{1, -2, 1}}
	if Intersection3DPlanePlane(p1, p2, got); !got.FuzzyEqual(want) {
		t.Error("Intersection3D.PlanePlane", *p1, *p2, "want", want, "got", got)
	}
}

func Benchmark_Intersection3D_PlanePlane(b *testing.B) {
	p1, p2 := &Plane{1, 1, 1, 1}, &Plane{1, 2, 3, 4}
	l := &Line3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DPlanePlane(p1, p2, l)
	}
}

func TestIntersection3DPlanePlanePlane(t *testing.T) {
	p1, p2, p3 := &Plane{1, -3, 3, 4}, &Plane{2, 3, -1, -15},
		&Plane{4, -3, -1, -19}
	got, want := &Vector3D{}, &Vector3D{5, 1, -2}
	if Intersection3DPlanePlanePlane(p1, p2, p3, got); !got.FuzzyEqual(want) {
		t.Error("Intersection3D.PlanePlanePlane", *p1, *p2, *p3, "want", want,
			"got", got)
	}
}

func Benchmark_Intersection3D_PlanePlanePlane(b *testing.B) {
	p1, p2, p3 := &Plane{1, -3, 3, 4}, &Plane{2, 3, -1, -15},
		&Plane{4, -3, -1, -19}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DPlanePlanePlane(p1, p2, p3, pt)
	}
}

type intersection3DFuzzyPlaneLineData struct {
	pl Plane
	l  Line3D
	pt Vector3D
	n  int
}

var intersection3DFuzzyPlaneLineValues = []intersection3DFuzzyPlaneLineData{
	{Plane{1, -2, 1, -7}, Line3D{Vector3D{5, -1, -1}, Vector3D{3, -1, 1}},
		Vector3D{11.0 / 2.0, -7.0 / 6.0, -5.0 / 6.0}, 1},
	{Plane{1, 0, 0, 0}, Line3D{Vector3D{1, 0, 0}, Vector3D{0, 1, 0}}, Vector3D{}, 0},
	{Plane{1, 0, 0, 0}, Line3D{Vector3D{}, Vector3D{0, 1, 0}}, Vector3D{}, -1},
}

func testIntersection3DFuzzyPlaneLine(d intersection3DFuzzyPlaneLineData, t *testing.T) {
	pt := &Vector3D{}
	got := Intersection3DFuzzyPlaneLine(&d.pl, &d.l, pt)
	if got != d.n {
		t.Error("Intersection3D.FuzzyPlaneLine", d.pl, d.l, "want", d.n, "got",
			got)
	} else if !d.pt.FuzzyEqual(pt) {
		t.Error("Intersection3D.FuzzyPlaneLine", d.pl, d.l, "want", d.pt,
			"got", *pt)
	}
}

func TestIntersection3DFuzzyPlaneLine(t *testing.T) {
	for _, v := range intersection3DFuzzyPlaneLineValues {
		testIntersection3DFuzzyPlaneLine(v, t)
		v.l.P.X, v.l.P.Y, v.l.P.Z = v.l.P.X+v.l.V.X, v.l.P.Y+v.l.V.Y, v.l.P.Z+v.l.V.Z
		v.l.V.X, v.l.V.Y, v.l.V.Z = -v.l.V.X, -v.l.V.Y, -v.l.V.Z
		testIntersection3DFuzzyPlaneLine(v, t)
	}
}

func Benchmark_Intersection3D_FuzzyPlaneLine(b *testing.B) {
	p := &Plane{1, 2, 3, 4}
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlaneLine(p, l, pt)
	}
}

type intersection3DFuzzyPlanePlaneData struct {
	p1, p2 Plane
	l      Line3D
	n      int
}

var intersection3DFuzzyPlanePlaneValue = []intersection3DFuzzyPlanePlaneData{
	{Plane{1, 0, 0, 1}, Plane{1, 0, 0, 1}, Line3D{Vector3D{}, Vector3D{}}, -1},
	{Plane{1, 0, 0, 1}, Plane{1, 0, 0, 2}, Line3D{Vector3D{}, Vector3D{}}, 0},
	{Plane{1, 1, 1, 1}, Plane{1, 2, 3, 4},
		Line3D{Vector3D{3, -5, 1}, Vector3D{4, -7, 2}}, 1},
	{Plane{1 + 1e-13, 0, 0, 1}, Plane{1, 0, 0, 1},
		Line3D{Vector3D{}, Vector3D{}}, -1},
}

func testIntersection3DFuzzyPlanePlane(d intersection3DFuzzyPlanePlaneData, t *testing.T) {
	var l Line3D
	if n := Intersection3DFuzzyPlanePlane(&d.p1, &d.p2, &l); n != d.n {
		t.Error("Intersection3D.FuzzyPlanePlane", d.p1, d.p2, "want", d.n,
			"got", n)
	}
}

func TestIntersection3DFuzzyPlanePlane(t *testing.T) {
	for _, v := range intersection3DFuzzyPlanePlaneValue {
		testIntersection3DFuzzyPlanePlane(v, t)
		v.p1, v.p2 = v.p2, v.p1
		testIntersection3DFuzzyPlanePlane(v, t)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlane(b *testing.B) {
	p1, p2 := &Plane{1, 2, 3, 4}, &Plane{5, 6, 7, 8}
	var l Line3D
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlane(p1, p2, &l)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlane_Coincident(b *testing.B) {
	p1, p2 := &Plane{1, 0, 0, 1}, &Plane{1, 0, 0, 1}
	var l Line3D
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlane(p1, p2, &l)
	}
}

type intersection3DFuzzyPlanePlanePlaneData struct {
	p1, p2, p3 Plane
	p          Vector3D
	n          int
}

var intersection3DFuzzyPlanePlanePlaneValues = []intersection3DFuzzyPlanePlanePlaneData{
	{Plane{1, -3, 3, 4}, Plane{2, 3, -1, -15}, Plane{4, -3, -1, -19},
		Vector3D{5, 1, -2}, 1},
	{Plane{1, 0, 0, 0}, Plane{1, 0, 0, 1}, Plane{1, 0, 0, 2}, Vector3D{}, 0},
	{Plane{1, 0, 0, 0}, Plane{1, 0, 0, 0}, Plane{1, 0, 0, 0}, Vector3D{}, -1},
	{Plane{1, 0, 0, 0}, Plane{0, 1, 0, 0}, Plane{1, 1, 0, 0}, Vector3D{}, -2},
	{Plane{1, 0, 0, 0}, Plane{1, 0, 0, 1}, Plane{1, 1, 0, 0}, Vector3D{}, -3},
	{Plane{-1, 0, 0, 0}, Plane{1, 0, 0, 1}, Plane{1, 1, 0, 0}, Vector3D{}, -3},
}

func testIntersection3DFuzzyPlanePlanePlane(d intersection3DFuzzyPlanePlanePlaneData, t *testing.T) {
	var p Vector3D
	n := Intersection3DFuzzyPlanePlanePlane(&d.p1, &d.p2, &d.p3, &p)
	if n != d.n {
		t.Error("Intersection3DFuzzyPlanePlanePlane", d.p1, d.p2, d.p3, "want",
			d.n, "got", n)
	} else if !p.FuzzyEqual(&d.p) {
		t.Error("Intersection3DFuzzyPlanePlanePlane", d.p1, d.p2, d.p3, "want",
			d.p, "got", p)
	}
}

func TestIntersection3DFuzzyPlanePlanePlane(t *testing.T) {
	for _, v := range intersection3DFuzzyPlanePlanePlaneValues {
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p2, v.p3 = v.p3, v.p2
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p1, v.p3 = v.p3, v.p1
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p2, v.p3 = v.p3, v.p2
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p1, v.p3 = v.p3, v.p1
		testIntersection3DFuzzyPlanePlanePlane(v, t)
		v.p2, v.p3 = v.p3, v.p2
		testIntersection3DFuzzyPlanePlanePlane(v, t)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlanePlane(b *testing.B) {
	p1, p2, p3 := &Plane{1, -3, 3, 4}, &Plane{2, 3, -1, -15},
		&Plane{4, -3, -1, -19}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlanePlane(p1, p2, p3, pt)
	}
}

func Benchmark_Intersection3D_FuzzyPlanePlanePlane_Line(b *testing.B) {
	p1, p2, p3 := &Plane{1, 0, 0, 0}, &Plane{0, 1, 0, 0}, &Plane{1, 1, 0, 0}
	pt := &Vector3D{}
	for i := 0; i < b.N; i++ {
		Intersection3DFuzzyPlanePlanePlane(p1, p2, p3, pt)
	}
}

type intersection3DRaySphereData struct {
	r Line3D
	s Sphere
	i Vector3D
	n int
}

var intersection3DRaySphereValues = []intersection3DRaySphereData{
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Sphere{Vector3D{}, 1}, Vector3D{1, 0, 0}, 1},
	{Line3D{Vector3D{0, 1, 0}, Vector3D{1, 0, 0}}, Sphere{Vector3D{}, 1}, Vector3D{0, 1, 0}, 1},
	{Line3D{Vector3D{0, 2, 0}, Vector3D{1, 0, 0}}, Sphere{Vector3D{}, 1}, Vector3D{0, 0, 0}, 0},
}

func testIntersection3DRaySphere(d intersection3DRaySphereData, t *testing.T) {
	var i Vector3D
	n := Intersection3DRaySphere(&d.r, &d.s, &i)
	if n != d.n {
		t.Error("Intersection3D.RaySphere", d.r, d.s, "want", d.n, "got", n)
		return
	}
	if n == 0 {
		return
	}
	if n == 1 && d.i.FuzzyEqual(&i) {
		return
	}
	t.Error("Intersection3D.RaySphere", d.r, d.s, "want", d.i, "got", i)
}

func TestIntersction3DRaySphere(t *testing.T) {
	for _, v := range intersection3DRaySphereValues {
		testIntersection3DRaySphere(v, t)
	}
}

func Benchmark_Intersection3D_RaySphere(b *testing.B) {
	l := Line3D{Vector3D{}, Vector3D{1, 0, 0}}
	s := Sphere{Vector3D{}, 1}
	var p Vector3D
	for i := 0; i < b.N; i++ {
		Intersection3DRaySphere(&l, &s, &p)
	}
}
