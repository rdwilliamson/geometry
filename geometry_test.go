package geometry

import "testing"

func TestFuzzyEqual(t *testing.T) {
	n1 := 0.0
	n2 := 1e-12
	n3 := 1e-13
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n2, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n3, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	n1 = -n1
	n2 = -n2
	n3 = -n3
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n2, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n3, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	n1 = 1e-12
	n2 = 0
	n3 = -1e-13
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n2, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n3, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	n1 = -n1
	n2 = -n2
	n3 = -n3
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n2, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n3, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	n1 = 1 + 1e-12
	n2 = 1 //- 1e-12
	n3 = 1 - 1e-13
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n2, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n3, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	n1 = 1e11
	n2 = 1e11 + 1e-1
	n3 = 1e11 + 1e-2
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n2, n1) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n3, n1) {
		t.Error("geometry.FuzzyEqual")
	}
}

func Benchmark_FuzzyEqual_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(1, 1+1e-13)
	}
}

func Benchmark_FuzzyEqual_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(-1, -1+1e-13)
	}
}

func Benchmark_FuzzyEqual_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(1+1e-13, 1)
	}
}

func Benchmark_FuzzyEqual_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(-1+1e-13, -1)
	}
}

func Benchmark_FuzzyEqual_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(-1e-13, 0)
	}
}

func Benchmark_FuzzyEqual_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(0, -1e-13)
	}
}

func Benchmark_FuzzyEqual_7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(1e-13, 0)
	}
}

func Benchmark_FuzzyEqual_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(0, 1e-13)
	}
}

func Benchmark_FuzzyEqual_All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(1+1e-13, 1)
		FuzzyEqual(-1, -1+1e-13)
		FuzzyEqual(1+1e-13, 1)
		FuzzyEqual(-1+1e-13, -1)
		FuzzyEqual(1e-13, 0)
		FuzzyEqual(0, 1e-13)
		FuzzyEqual(-1e-13, 0)
		FuzzyEqual(0, -1e-13)
	}
}
