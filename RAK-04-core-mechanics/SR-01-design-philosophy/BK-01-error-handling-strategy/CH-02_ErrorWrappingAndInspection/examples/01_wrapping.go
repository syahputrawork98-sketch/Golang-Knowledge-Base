package main

import (
	"errors"
	"fmt"
)

// 01_wrapping.go - Membungkus error (Error Context Wrapping)
// Analogi: Memasukkan paket rusak ke dalam kotak baru dengan label tambahan.

var ErrDatabaseLost = errors.New("koneksi database terputus")

func fetchData() error {
	// Membungkus error asli dengan konteks tambahan menggunakan %w
	return fmt.Errorf("gagal mengambil data user: %w", ErrDatabaseLost)
}

func main() {
	err := fetchData()
	if err != nil {
		fmt.Printf("Log Error Sistem: %v\n", err)
		
		// Kita tetap bisa mengakses error aslinya (Unwrap internal)
		// errors.Is akan menelusuri rantai wrapping
		if errors.Is(err, ErrDatabaseLost) {
			fmt.Println("Tindakan: Melakukan Reconnect Otomatis...")
		}
	}
}
