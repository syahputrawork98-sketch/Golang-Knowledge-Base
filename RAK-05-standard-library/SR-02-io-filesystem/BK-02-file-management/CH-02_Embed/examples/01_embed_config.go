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
	fmt.Printf("Key ditemukan dari dalam biner: %s\n", secretKey)
}
