package bench

import "testing"

type SmallT struct {
	X int32 // 4B
}

type HugeT struct {
	X [1000]int32 // 4KB
}

type SuperHugeT struct {
	X [10 * 1000 * 1000]byte // 10MB
}

var (
	smallG     interface{}
	hugeG      interface{}
	superHugeG interface{}
)

func BenchmarkSmallAllocOnHeap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallG = &SmallT{}
	}
}

func BenchmarkSmallAllocOnStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		local := SmallT{}
		_ = local
	}
}

func BenchmarkHugeAllocOnHeap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hugeG = &HugeT{}
	}
}

func BenchmarkHugeAllocOnStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		local := HugeT{}
		_ = local
	}
}

func BenchmarkSuperHugeAllocOnHeap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		superHugeG = &SuperHugeT{}
	}
}

func BenchmarkSuperHugeAllocOnStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		local := SuperHugeT{}
		_ = local
	}
}
