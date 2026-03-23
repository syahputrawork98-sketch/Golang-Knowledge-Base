package main

import "fmt"

// Go menggunakan Dictionary untuk mengirim informasi tipe
// ke fungsi generic. Ini memiliki sedikit biaya overhead.

func PrintAnything[T any](v T) {
	fmt.Printf("Val: %v, Type: %T\n", v, v)
}

func main() {
	// Panggilan ini akan menyertakan dictionary pointer di balik layar
	PrintAnything(42)
	PrintAnything("Generics")
	
	fmt.Println("\nInfo: Meskipun efisien, penggunaan generics berlebihan pada fungsi " +
		"sederhana yang bisa memakai 'any' mungkin tidak selalu memberikan keuntungan performa signifikan.")
}
