package reverse

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	// 1. Tambahkan Seed Corpus (Input awal)
	f.Add("Hello")
	f.Add("世界")
	f.Add("!@#")

	// 2. Jalankan Fuzz Engine
	f.Fuzz(func(t *testing.T, orig string) {
		// 3. Tes Invarian: Dobel reverse harus kembali ke asli
		rev, err := Reverse(orig)
		if err != nil {
			return
		}

		doubleRev, err := Reverse(rev)
		if err != nil {
			return
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, after double reverse: %q", orig, doubleRev)
		}

		// 4. Tes Invarian UTF-8
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse of valid UTF-8 string is not valid: %q", rev)
		}
	})
}
