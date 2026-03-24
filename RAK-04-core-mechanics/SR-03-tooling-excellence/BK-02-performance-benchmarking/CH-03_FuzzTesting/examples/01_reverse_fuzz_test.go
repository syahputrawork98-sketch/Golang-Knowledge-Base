package examples

import (
	"testing"
	"unicode/utf8"
)

// 01_reverse_fuzz_test.go - Mengetes Jembatan (Stress Test)
// Jalankan dengan: go test -fuzz=FuzzReverse

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello", " ", "!123"}
	for _, tc := range testcases {
		f.Add(tc) // Seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, After: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
