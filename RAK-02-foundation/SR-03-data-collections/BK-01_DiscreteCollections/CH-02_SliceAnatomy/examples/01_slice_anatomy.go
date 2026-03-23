package main

import "fmt"

func main() {
	// 1. Membuat slice dengan make: make([]T, len, cap)
	s := make([]int, 3, 5)

	fmt.Printf("Slice s: %v | Len: %d | Cap: %d\n", s, len(s), cap(s))

	// 2. Mengisi data
	s[0], s[1], s[2] = 10, 20, 30
	fmt.Println("Setelah diisi:", s)

	// 3. Menambah data melewati len tapi dalam batas cap (Gunakan append)
	s = append(s, 40)
	fmt.Printf("Setelah append: %v | Len: %d | Cap: %d\n", s, len(s), cap(s))

	// Kesimpulan: Length adalah apa yang bisa kita akses, 
	// Capacity adalah seberapa besar backing array sebelum perlu dialokasi ulang.
}
