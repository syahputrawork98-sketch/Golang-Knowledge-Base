package main

import "fmt"

type DataFetcher interface {
	FetchData() string
}

// Implementasi Asli (Lambat/Butuh Koneksi)
type RealAPI struct{}
func (r RealAPI) FetchData() string { return "Real Data from Cloud" }

// Implementasi Mock (Cepat/Untuk Testing)
type MockAPI struct{}
func (m MockAPI) FetchData() string { return "Mocked Test Data" }

func ProcessData(f DataFetcher) {
	data := f.FetchData()
	fmt.Println("Processing:", data)
}

func main() {
	// Di Production:
	real := RealAPI{}
	ProcessData(real)

	// Di Testing:
	mock := MockAPI{}
	ProcessData(mock)
	
	fmt.Println("\nKesimpulan: Interface memungkinkan penggantian komponen secara instan.")
}
