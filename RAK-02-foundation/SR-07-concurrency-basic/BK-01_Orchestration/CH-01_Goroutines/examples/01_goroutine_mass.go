package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Ambil statistik memori sebelum mulai
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	
	count := 100000 // 100 Ribu Goroutines
	
	fmt.Printf("Mencoba menjalankan %d Goroutines...\n", count)
	
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(10 * time.Second) // Menahan goroutine agar tidak langsung selesai
		}()
	}

	time.Sleep(2 * time.Second) // Tunggu agar semua teralokasi

	// Ambil statistik memori setelah mulai
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("Memory Used: %d KB\n", (m2.Alloc - m1.Alloc) / 1024)
	fmt.Printf("Rata-rata per Goroutine: %.2f bytes\n", float64(m2.Alloc-m1.Alloc)/float64(count))
	
	fmt.Println("\nKesimpulan: Mengejutkan bukan? 100 ribu tugas hanya memakan sedikit memori!")
}
