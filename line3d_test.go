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

type line3DFuzzyEqualData struct {
	l1, l2 Line3D
	equal  bool
}

var line3DFuzzyEqualValues = []line3DFuzzyEqualData{
	{Line3D{Vector3D{}, Vector3D{1, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1, 1, 1 + 1e-12}}, false},
	{Line3D{Vector3D{}, Vector3D{1, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1, 1, 1 + 1e-13}}, true},
	{Line3D{Vector3D{}, Vector3D{1, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1, 1 + 1e-12, 1}}, false},
	{Line3D{Vector3D{}, Vector3D{1, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1, 1 + 1e-13, 1}}, true},
	{Line3D{Vector3D{}, Vector3D{1, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1 + 1e-12, 1, 1}}, false},
	{Line3D{Vector3D{}, Vector3D{1, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1 + 1e-13, 1, 1}}, true},
	{Line3D{Vector3D{1, 0, 0}, Vector3D{2, 1, 1}},
		Line3D{Vector3D{}, Vector3D{1, 1, 1}}, false},
}

func testLine3DFuzzyEqual(d line3DFuzzyEqualData, t *testing.T) {
	if d.l1.FuzzyEqual(&d.l2) != d.equal {
		t.Error("Line3D.FuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.FuzzyEqual(&d.l1) != d.equal {
		t.Error("Line3D.FuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if d.l1.FuzzyEqual(&d.l2) != d.equal {
		t.Error("Line3D.FuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.FuzzyEqual(&d.l1) != d.equal {
		t.Error("Line3D.FuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if d.l1.FuzzyEqual(&d.l2) != d.equal {
		t.Error("Line3D.FuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.FuzzyEqual(&d.l1) != d.equal {
		t.Error("Line3D.FuzzyEqual", d.l2, d.l1, d.equal)
	}
}

func TestLine3DFuzzyEqual(t *testing.T) {
	for _, v := range line3DFuzzyEqualValues {
		testLine3DFuzzyEqual(v, t)
	}
}

func Benchmark_Line3D_FuzzyEqual(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}}
	l2 := &Line3D{Vector3D{5, 6, 7}, Vector3D{0, 1, 2}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
	}
}

type line3DLineBetweenData struct {
	l1, l2, lb Line3D
}

var line3DLineBetweenValues = []line3DLineBetweenData{
	{Line3D{Vector3D{1, 2, 1}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}},
		Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}},
	{Line3D{Vector3D{5, 4, 5}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, 5}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{-3, 0, -3}},
		Line3D{Vector3D{5, 6, 1}, Vector3D{9, 8, -3}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}},
}

func testLine3DLineBetween(d line3DLineBetweenData, t *testing.T) {
	var l Line3D
	if !l.LineBetween(&d.l1, &d.l2).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.LineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
	if !l.LineBetween(&d.l2, &d.l1).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.LineBetween", d.l2, d.l1, "want", d.lb, "got", l)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if !l.LineBetween(&d.l1, &d.l2).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.LineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
	if !l.LineBetween(&d.l2, &d.l1).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.LineBetween", d.l2, d.l1, "want", d.lb, "got", l)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if !l.LineBetween(&d.l1, &d.l2).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.LineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
	if !l.LineBetween(&d.l2, &d.l1).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.LineBetween", d.l2, d.l1, "want", d.lb, "got", l)
	}
}

func TestLine3DLineBetween(t *testing.T) {
	for _, v := range line3DLineBetweenValues {
		testLine3DLineBetween(v, t)
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

func TestLine3DMidpoint(t *testing.T) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{6, 5, 4}}
	p := &Vector3D{}
	if !l.Midpoint(p).Equal(&Vector3D{3.5, 3.5, 3.5}) {
		t.Error("Line3D.Midpoint")
	}
}

func Benchmark_Line3D_Midpoint(b *testing.B) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{6, 5, 4}}
	p := &Vector3D{}
	for i := 0; i < b.N; i++ {
		l.Midpoint(p)
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

type line3DSegmentFuzzyEqualData struct {
	l1, l2 Line3D
	equal  bool
}

var line3DSegmentFuzzyEqualValues = []line3DSegmentFuzzyEqualData{
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6 + 1e-11}}, false},
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6 + 1e-12}}, true},
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5 + 1e-11, 6}}, false},
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5 + 1e-12, 6}}, true},
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{1, 2, 3}, Vector3D{4 + 1e-11, 5, 6}}, false},
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{1, 2, 3}, Vector3D{4 + 1e-12, 5, 6}}, true},
	{Line3D{Vector3D{1, 2, 3}, Vector3D{4, 5, 6}},
		Line3D{Vector3D{2, 2, 3}, Vector3D{5, 5, 6}}, false},
}

func testLine3DSegmentFuzzyEqualData(d line3DSegmentFuzzyEqualData, t *testing.T) {
	if d.l1.SegmentFuzzyEqual(&d.l2) != d.equal {
		t.Error("Line3D.SegmentFuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.SegmentFuzzyEqual(&d.l1) != d.equal {
		t.Error("Line3D.SegmentFuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if d.l1.SegmentFuzzyEqual(&d.l2) != d.equal {
		t.Error("Line3D.SegmentFuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.SegmentFuzzyEqual(&d.l1) != d.equal {
		t.Error("Line3D.SegmentFuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if d.l1.SegmentFuzzyEqual(&d.l2) != d.equal {
		t.Error("Line3D.SegmentFuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.SegmentFuzzyEqual(&d.l1) != d.equal {
		t.Error("Line3D.SegmentFuzzyEqual", d.l2, d.l1, d.equal)
	}
}

func TestLine3DSegmentFuzzyEqual(t *testing.T) {
	for _, v := range line3DSegmentFuzzyEqualValues {
		testLine3DSegmentFuzzyEqualData(v, t)
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

func TestLine3DPointDistanceSquared(t *testing.T) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	if l.PointDistanceSquared(p) != 1 {
		t.Error("Line3D.PointDistanceSquared")
	}
}

func Benchmark_Line3D_PointDistanceSquared(b *testing.B) {
	l := Line3D{Vector3D{0, 0, 0}, Vector3D{1, 0, 0}}
	p := &Vector3D{0, 1, 0}
	for i := 0; i < b.N; i++ {
		l.PointDistanceSquared(p)
	}
}

type line3DSegmentLineBetweenData struct {
	l1, l2, lb Line3D
}

var line3DSegmentLineBetweenValues = []line3DSegmentLineBetweenData{
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{-3, 0, -3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, 3}},
		Line3D{Vector3D{-1, 1, -1}, Vector3D{1, 2, 1}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{3, 3, 3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, -3}},
		Line3D{Vector3D{1, 2, 1}, Vector3D{1, 2, -1}}},
	{Line3D{Vector3D{-1, 1, -1}, Vector3D{-3, 0, -3}},
		Line3D{Vector3D{1, 2, -1}, Vector3D{1, 2, -3}},
		Line3D{Vector3D{-1, 1, -1}, Vector3D{1, 2, -1}}},
	{Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}},
		Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}},
		Line3D{Vector3D{3.4, 3.2, 3.4}, Vector3D{2.6, 4.8, 3.4}}},
}

func testLine3DSegmentLineBetween(d line3DSegmentLineBetweenData, t *testing.T) {
	var l Line3D
	if !l.SegmentLineBetween(&d.l1, &d.l2).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.SegmentLineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
	if !l.SegmentLineBetween(&d.l2, &d.l1).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.SegmentLineBetween", d.l2, d.l1, "want", d.lb, "got", l)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if !l.SegmentLineBetween(&d.l1, &d.l2).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.SegmentLineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
	if !l.SegmentLineBetween(&d.l2, &d.l1).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.SegmentLineBetween", d.l2, d.l1, "want", d.lb, "got", l)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if !l.SegmentLineBetween(&d.l1, &d.l2).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.SegmentLineBetween", d.l1, d.l2, "want", d.lb, "got", l)
	}
	if !l.SegmentLineBetween(&d.l2, &d.l1).SegmentFuzzyEqual(&d.lb) {
		t.Error("Line3D.SegmentLineBetween", d.l2, d.l1, "want", d.lb, "got", l)
	}
}

func TestLine3DSegmentLineBetween(t *testing.T) {
	for _, v := range line3DSegmentLineBetweenValues {
		testLine3DSegmentLineBetween(v, t)
	}
}

func Benchmark_Line3D_SegmentLineBetween(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 1}, Vector3D{5, 4, 5}}
	l2 := &Line3D{Vector3D{5, 6, 1}, Vector3D{1, 4, 5}}
	r := &Line3D{}
	for i := 0; i < b.N; i++ {
		r.SegmentLineBetween(l1, l2)
	}
}

type line3DSegmentPointDistanceData struct {
	l Line3D
	p Vector3D
	d float64
}

var line3DSegmentPointDistanceValues = []line3DSegmentPointDistanceData{
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{0, 2, 0}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{0, 0, 2}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{3, 0, 0}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{-2, 0, 0}, 2},
	{Line3D{Vector3D{}, Vector3D{1, 0, 0}}, Vector3D{0.5, 0, 0}, 0},
}

func testLine3DSegmentPointDistance(d line3DSegmentPointDistanceData, t *testing.T) {
	if got := d.l.SegmentPointDistance(&d.p); got != d.d {
		t.Error("Line3D.SegmentPointDistance", d.l, d.p, "want", d.d, "got", got)
	}
}

func TestLine3DSegmentPointDistance(t *testing.T) {
	for _, v := range line3DSegmentPointDistanceValues {
		testLine3DSegmentPointDistance(v, t)
	}
}

func Benchmark_Line3D_SegmentPointDistance(b *testing.B) {
	l := &Line3D{Vector3D{}, Vector3D{1, 1, 1}}
	p := &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistance(p)
	}
}

func testLine3DSegmentPointDistanceSquared(d line3DSegmentPointDistanceData, t *testing.T) {
	if got := d.l.SegmentPointDistanceSquared(&d.p); got != d.d*d.d {
		t.Error("Line3D.SegmentPointDistanceSquared", d.l, d.p, "want", d.d*d.d, "got", got)
	}
}

func TestLine3DSegmentPointDistanceSquared(t *testing.T) {
	for _, v := range line3DSegmentPointDistanceValues {
		testLine3DSegmentPointDistanceSquared(v, t)
	}
}

func Benchmark_Line3D_SegmentPointDistanceSquared(b *testing.B) {
	l := &Line3D{Vector3D{}, Vector3D{1, 1, 1}}
	p := &Vector3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		l.SegmentPointDistanceSquared(p)
	}
}

func TestLine3DSet(t *testing.T) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{}, Vector3D{}}
	if !l2.Set(l1).Equal(l1) {
		t.Error("Line3D.Set")
	}
}

func Benchmark_Line3D_Set(b *testing.B) {
	l1 := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	l2 := &Line3D{Vector3D{}, Vector3D{}}
	for i := 0; i < b.N; i++ {
		l2.Set(l1)
	}
}

func TestLine3DToVector(t *testing.T) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	v := &Vector3D{}
	if !l.ToVector(v).Equal(&Vector3D{0, -1, -2}) {
		t.Error("Line3D.ToVector")
	}
}

func Benchmark_Line3D_ToVector(b *testing.B) {
	l := &Line3D{Vector3D{1, 2, 3}, Vector3D{1, 1, 1}}
	v := &Vector3D{}
	for i := 0; i < b.N; i++ {
		l.ToVector(v)
	}
}
