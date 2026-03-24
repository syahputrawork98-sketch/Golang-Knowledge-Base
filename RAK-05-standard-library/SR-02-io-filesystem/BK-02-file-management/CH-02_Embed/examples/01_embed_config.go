package main

import (
	_ "embed"
	"fmt"
)

// 01_embed_config.go - Menyisipkan File ke Biner
// Analogi: Mencetak foto langsung di dalam buku (melekat).

//go:embed test_data.txt
var secretKey string

func main() {
	// Syarat: File 'test_data.txt' harus ada di folder yang sama saat build
	// Kita simulasi pembacaan jika file berhasil di-embed.
	
	if secretKey == "" {
		fmt.Println("Hint: Jalankan 'echo 'KEY_123' > test_data.txt' lalu build.")
	} else {
		fmt.Printf("Key ditemukan dari dalam biner: %s\n", secretKey)
	}
}
