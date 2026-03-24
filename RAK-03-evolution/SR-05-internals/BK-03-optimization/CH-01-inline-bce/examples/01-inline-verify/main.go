package main

import "fmt"

// smallFunc adalah kandidat utama inlining karena sederhana
func smallFunc(a, b int) int {
	return a + b
}

// complexFunc mungkin tidak di-inline karena terlalu besar (misal ada loop)
func complexFunc(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

func main() {
	// 1. Verifikasi Inlining
	x := smallFunc(10, 20)
	y := complexFunc(x)
	fmt.Println(y)

	// 2. Verifikasi Bounds Check Elimination (BCE)
	data := []int{1, 2, 3, 4, 5}
	
	// Trik BCE: Cek indeks terakhir di awal
	// Compiler akan tahu bahwa index 0-4 aman di baris berikutnya.
	if len(data) >= 5 {
		_ = data[4] // Pengecekan dilakukan di sini
		
		fmt.Println(data[0]) // BCE: Tidak ada pengecekan batas di sini
		fmt.Println(data[1]) // BCE: Tidak ada pengecekan batas di sini
		fmt.Println(data[2]) // BCE: Tidak ada pengecekan batas di sini
	}

	// Lab Guide:
	// Jalankan perintah berikut untuk melihat laporan compiler:
	// go build -gcflags="-m -l" main.go
	// -> -m : Print optimization decisions
	// -> -l : Disable inlining (opsional, untuk perbandingan)
}
