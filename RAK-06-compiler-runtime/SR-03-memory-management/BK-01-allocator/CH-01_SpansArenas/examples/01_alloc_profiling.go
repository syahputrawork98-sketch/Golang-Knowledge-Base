package main

import (
	"fmt"
	"runtime"
	"time"
)

// 01_alloc_profiling.go - Profiling Alokasi Memori
// Analogi: Memasang CCTV di gudang untuk tahu barang mana yang menghabiskan tempat.

func main() {
	// Jalankan program dengan profiling memori aktif
	// go run 01_alloc_profiling.go
	
	fmt.Println("Starting allocation marathon...")
	
	// Simulasi alokasi yang sering terjadi (berpotensi escape to heap)
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1024) // Alokasi 1KB per iterasi
		if i % 10000 == 0 {
			printStats()
		}
	}

	fmt.Println("Allocation Finished.")
	time.Sleep(1 * time.Second)
}

func printStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
