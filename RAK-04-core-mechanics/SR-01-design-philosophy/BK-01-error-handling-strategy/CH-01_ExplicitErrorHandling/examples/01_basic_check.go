package main

import (
	"fmt"
	"os"
)

// 01_basic_check.go - Demonstrasi pola idiomatik "if err != nil"
// Analogi: Kurir yang memberikan laporan tertulis jika pengiriman gagal.

func main() {
	// Membuka file yang mungkin tidak ada
	file, err := os.Open("data_rahasia.txt")
	
	// Tahap 1: Pengecekan Eksplisit (The Art of Checking)
	if err != nil {
		// Logika penanganan error diletakkan sedini mungkin
		fmt.Printf("Error ditemukan: %v\n", err)
		return
	}
	
	// Tahap 2: Happy Path (Tanpa Indentasi Berlebih)
	defer file.Close()
	fmt.Println("File berhasil dibuka!")
}
