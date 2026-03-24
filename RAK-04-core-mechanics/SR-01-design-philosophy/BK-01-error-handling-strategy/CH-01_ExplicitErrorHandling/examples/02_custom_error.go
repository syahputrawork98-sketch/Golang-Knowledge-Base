package main

import (
	"fmt"
)

// 02_custom_error.go - Membuat tipe error kustom (The Expert Case)
// Analogi: Laporan Kurir yang mendetail (kehabisan bensin, ban bocor, dsb).

// RequestError menyimpan informasi tambahan tentang kegagalan API
type RequestError struct {
	StatusCode int
	Message    string
}

// Error() memenuhi interface error bawaan Go
func (r *RequestError) Error() string {
	return fmt.Sprintf("API failure [%d]: %s", r.StatusCode, r.Message)
}

func callExternalAPI() error {
	// Simulasi kegagalan
	return &RequestError{
		StatusCode: 404,
		Message:    "Data personil tidak ditemukan",
	}
}

func main() {
	err := callExternalAPI()
	if err != nil {
		// Kita bisa menggunakan type assertion untuk mengecek data detail
		if reqErr, ok := err.(*RequestError); ok {
			fmt.Printf("Detail Error - Kode: %d, Pesan: %s\n", reqErr.StatusCode, reqErr.Message)
		} else {
			fmt.Printf("General Error: %v\n", err)
		}
	}
}
