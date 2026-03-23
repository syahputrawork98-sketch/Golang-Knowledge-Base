package main

import "fmt"

func main() {
	// 1. Inisialisasi dengan make
	scores := make(map[string]int)

	// 2. Menambah & Update data
	scores["Alice"] = 90
	scores["Bob"] = 85
	scores["Alice"] = 95 // Override

	// 3. Membaca data dengan Comma OK Idiom
	val, ok := scores["Charlie"]
	if !ok {
		fmt.Println("Charlie tidak ditemukan!")
	} else {
		fmt.Println("Skor Charlie:", val)
	}

	// 4. Menghapus data
	delete(scores, "Bob")

	fmt.Println("Final Map:", scores)
	fmt.Println("Panjang Map:", len(scores))
}
