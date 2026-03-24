package main

import (
	"fmt"
	"runtime"
	"time"
)

// 01_gmp_inspection.go - Mengintip Isi Scheduler
// Analogi: Melihat dashboard dapur untuk tahu berapa koki (M) dan meja (P) yang aktif.

func main() {
	// 1. Melihat jumlah Processor (P) default (sesuai Core CPU)
	fmt.Printf("GOMAXPROCS (P): %d\n", runtime.GOMAXPROCS(0))

	// 2. Monitoring jumlah Goroutine (G)
	fmt.Printf("Initial Goroutines: %d\n", runtime.NumGoroutine())

	// 3. Menjalankan ribuan G secara paralel
	for i := 0; i < 1000; i++ {
		go func(id int) {
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	fmt.Printf("Active Goroutines (after spawning 1000): %d\n", runtime.NumGoroutine())
	
	// Memberi waktu scheduler bekerja
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Final Goroutines: %d\n", runtime.NumGoroutine())
}
