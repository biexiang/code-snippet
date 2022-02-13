package inline_optimize

import "testing"

//go:noinline
func maxNoInline(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxInline(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func BenchmarkNoInline(b *testing.B) {
	x, y := 1, 2
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = maxNoInline(x, y)
	}
}

func BenchmarkInline(b *testing.B) {
	x, y := 1, 2
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = maxInline(x, y)
	}
}
