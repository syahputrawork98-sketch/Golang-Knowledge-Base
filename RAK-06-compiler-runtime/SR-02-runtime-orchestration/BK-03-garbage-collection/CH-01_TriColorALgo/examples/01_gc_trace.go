package main

import (
	"fmt"
	"runtime"
	"time"
)

// 01_gc_trace.go - Memantau Kolektor Sampah
// Analogi: Memanggil petugas kebersihan secara manual dan melihat laporannya.

func main() {
	var m runtime.MemStats

	// 1. Alokasi memori besar
	data := make([]byte, 50*1024*1024) // 50MB
	_ = data

	runtime.ReadMemStats(&m)
	fmt.Printf("Allocated: %d MB\n", m.Alloc/1024/1024)

	// 2. Trigger GC secara manual
	fmt.Println("Triggering GC...")
	runtime.GC()

	// 3. Lihat statistik setelah GC
	runtime.ReadMemStats(&m)
	fmt.Printf("Allocated after GC: %d MB\n", m.Alloc/1024/1024)
	fmt.Printf("Num GC: %d\n", m.NumGC)

	// TIPS: Jalankan dengan GODEBUG=gctrace=1 go run 01_gc_trace.go
	// untuk melihat detail tri-color marking di terminal.
	time.Sleep(1 * time.Second)
}
