package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Kata "Hello" (ASCII) dan "世界" (Chinese - Hello)
	s1 := "Hello"
	s2 := "世界"

	fmt.Println("--- String Internals: Byte vs Rune ---")

	// 1. Menggunakan len() standar (menghitung byte)
	fmt.Printf("String: '%s', Bytes: %d\n", s1, len(s1))
	fmt.Printf("String: '%s', Bytes: %d (Satu karakter Chinese = 3 byte!)\n", s2, len(s2))

	// 2. Menghitung Rune (Karakter Unicode)
	fmt.Printf("String: '%s', Runes: %d\n", s2, utf8.RuneCountInString(s2))

	// 3. Iterasi otomatis menjadi rune
	fmt.Println("\nIterating over '世界':")
	for i, r := range s2 {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}
}
