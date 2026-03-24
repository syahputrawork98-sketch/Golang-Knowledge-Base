package examples

import (
	"strings"
	"testing"
)

// 01_string_concat_test.go - Benchmark Kecepatan (F1 Race)
// Jalankan dengan: go test -bench=. 01_string_concat_test.go

func BenchmarkPlusConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "Go"
		}
	}
}

func BenchmarkBuilderConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 100; j++ {
			builder.WriteString("Go")
		}
		_ = builder.String()
	}
}
