package main

import (
	"fmt"
	"strings"
)

// 02_string_builder.go - Efisiensi Alokasi Memori
// Analogi: Menulis di papan tulis besar sekaligus, bukan mencetak kertas baru setiap kata.

func main() {
	// strings.Builder meminimalkan penyalinan memori saat penggabungan
	var builder strings.Builder

	words := []string{"Membangun", "masa", "depan", "dengan", "Go", "Standard", "Library."}

	for _, word := range words {
		// Menambahkan ke buffer internal builder
		builder.WriteString(word)
		builder.WriteString(" ")
	}

	// Mengambil hasil akhir sebagai satu string utuh
	result := builder.String()
	fmt.Println("Result:", result)
	fmt.Printf("Builder size: %d bytes\n", builder.Len())
}
