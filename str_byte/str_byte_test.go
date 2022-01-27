package str_byte

import (
	"testing"

	"github.com/biexiang/code-snippet/str_byte/normal"
	"github.com/biexiang/code-snippet/str_byte/unsafe"
)

func BenchmarkStrByte(b *testing.B) {

	s := "may the force be with you !"
	bt := []byte(s)

	b.ResetTimer()
	b.Run("NormalBenchmarkStr2Byte", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = normal.Str2byte(s)
		}
	})

	b.ResetTimer()
	b.Run("UnsafeBenchmarkStr2Byte", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = unsafe.Str2byte(s)
		}
	})

	b.ResetTimer()
	b.Run("NormalBenchmarkByte2Str", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = normal.Byte2str(bt)
		}
	})

	b.ResetTimer()
	b.Run("UnsafeBenchmarkByte2Str", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = unsafe.Byte2str(bt)
		}
	})

}
