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

func TestLine2DCopy(t *testing.T) {
	l1, l2 := &Line2D{}, &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	if !l1.Copy(l2).SegmentEqual(l2) {
		t.Error("Line2D.Copy")
	}
}

func Benchmark_Line2D_Copy(b *testing.B) {
	l1, l2 := &Line2D{}, &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l1.Copy(l2)
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

type line2DFuzzyEqualData struct {
	l1, l2 Line2D
	equal  bool
}

var line2DFuzzyEqualValues = []line2DFuzzyEqualData{
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{-3, -2}, Vector2D{5, 6 + 1e-11}}, false},
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{-3, -2}, Vector2D{5, 6 + 1e-12}}, true},
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{-3, -2}, Vector2D{5 + 1e-11, 6}}, false},
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{-3, -2}, Vector2D{5 + 1e-12, 6}}, true},
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{2, 2}, Vector2D{4, 4}}, false},
}

func testLine2DFuzzyEqual(d line2DFuzzyEqualData, t *testing.T) {
	if d.l1.FuzzyEqual(&d.l2) != d.equal {
		t.Error("Line2D.FuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.FuzzyEqual(&d.l1) != d.equal {
		t.Error("Line2D.FuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if d.l1.FuzzyEqual(&d.l2) != d.equal {
		t.Error("Line2D.FuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.FuzzyEqual(&d.l1) != d.equal {
		t.Error("Line2D.FuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if d.l1.FuzzyEqual(&d.l2) != d.equal {
		t.Error("Line2D.FuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.FuzzyEqual(&d.l1) != d.equal {
		t.Error("Line2D.FuzzyEqual", d.l2, d.l1, d.equal)
	}
}

func TestLine2DFuzzyEqual(t *testing.T) {
	for _, v := range line2DFuzzyEqualValues {
		testLine2DFuzzyEqual(v, t)
	}
}

func Benchmark_Line2D_FuzzyEqual(b *testing.B) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{-3, -2}, Vector2D{5, 6}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
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

func TestLine2DSegmentEqual(t *testing.T) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{3, 4}, Vector2D{1, 2}}
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

type line2DSegmentFuzzyEqualData struct {
	l1, l2 Line2D
	equal  bool
}

var line2DSegmentFuzzyEqualValues = []line2DSegmentFuzzyEqualData{
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{1, 2}, Vector2D{3, 4 + 1e-11}}, false},
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{1, 2}, Vector2D{3, 4 + 1e-12}}, true},
	{Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, Line2D{Vector2D{2, 3}, Vector2D{4, 5}}, false},
}

func testLine2DSegmentFuzzyEqual(d line2DSegmentFuzzyEqualData, t *testing.T) {
	if d.l1.SegmentFuzzyEqual(&d.l2) != d.equal {
		t.Error("Line2D.SegmentFuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.SegmentFuzzyEqual(&d.l1) != d.equal {
		t.Error("Line2D.SegmentFuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l1.P1, d.l1.P2 = d.l1.P2, d.l1.P1
	if d.l1.SegmentFuzzyEqual(&d.l2) != d.equal {
		t.Error("Line2D.SegmentFuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.SegmentFuzzyEqual(&d.l1) != d.equal {
		t.Error("Line2D.SegmentFuzzyEqual", d.l2, d.l1, d.equal)
	}
	d.l2.P1, d.l2.P2 = d.l2.P2, d.l2.P1
	if d.l1.SegmentFuzzyEqual(&d.l2) != d.equal {
		t.Error("Line2D.SegmentFuzzyEqual", d.l1, d.l2, d.equal)
	}
	if d.l2.SegmentFuzzyEqual(&d.l1) != d.equal {
		t.Error("Line2D.SegmentFuzzyEqual", d.l2, d.l1, d.equal)
	}
}

func TestLine2DSegmentFuzzyEqual(t *testing.T) {
	for _, v := range line2DSegmentFuzzyEqualValues {
		testLine2DSegmentFuzzyEqual(v, t)
	}
}

func Benchmark_Line2D_SegmentFuzzyEqual(b *testing.B) {
	l1 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	l2 := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l1.SegmentFuzzyEqual(l2)
	}
}

type line2DSlopeData struct {
	l     Line2D
	slope float64
}

var line2DSlopeValues = []line2DSlopeData{
	{Line2D{Vector2D{}, Vector2D{1, 1}}, 1},
	{Line2D{Vector2D{}, Vector2D{1, -1}}, -1},
	{Line2D{Vector2D{}, Vector2D{-1, 1}}, -1},
	{Line2D{Vector2D{}, Vector2D{-1, -1}}, 1},
	{Line2D{Vector2D{}, Vector2D{1, 0}}, 0},
	{Line2D{Vector2D{}, Vector2D{0, 1}}, math.Inf(1)},
	{Line2D{Vector2D{}, Vector2D{0, -1}}, math.Inf(-1)},
}

func testLine2DSlope(d line2DSlopeData, t *testing.T) {
	if d.l.Slope() != d.slope {
		t.Error("Line2D.Slope", d.l, d.slope)
	}
}

func TestLine2DSlope(t *testing.T) {
	for _, v := range line2DSlopeValues {
		testLine2DSlope(v, t)
	}
}

func Benchmark_Line2D_Slope(b *testing.B) {
	l := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}
	for i := 0; i < b.N; i++ {
		l.Slope()
	}
}

func TestLine2DToVector(t *testing.T) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	if !l.ToVector(v).Equal(&Vector2D{2, 2}) {
		t.Error("Line2D.ToVector")
	}
}

func Benchmark_Vector2D_ToVector(b *testing.B) {
	l, v := &Line2D{Vector2D{1, 2}, Vector2D{3, 4}}, &Vector2D{}
	for i := 0; i < b.N; i++ {
		l.ToVector(v)
	}
}
