package main

import (
	"fmt"
	"runtime"
	"time"
)

// simulateAllocations memicu pembuatan banyak span baru
func simulateAllocations() {
	var holder [][]byte
	for i := 0; i < 10000; i++ {
		// Alokasi objek kecil (Size Class rendah)
		b := make([]byte, 64)
		holder = append(holder, b)
	}
}

func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", label)
	fmt.Printf("Alloc      = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys        = %v MiB (Memori dari OS)\n", m.Sys/1024/1024)
	fmt.Printf("NumGC      = %v\n", m.NumGC)
	fmt.Printf("MSpanInuse = %v KiB\n", m.MSpanInuse/1024)
	fmt.Println("-------------------------------")
}

func main() {
	printMemStats("Status Awal")

	fmt.Println("Simulating internal mspan allocations...")
	simulateAllocations()
	
	printMemStats("Setelah Alokasi")

	fmt.Println("Triggering manual GC & sweeping...")
	runtime.GC()
	time.Sleep(100 * time.Millisecond) // Tunggu sweeper selesai
	
	printMemStats("Setelah GC")
}
