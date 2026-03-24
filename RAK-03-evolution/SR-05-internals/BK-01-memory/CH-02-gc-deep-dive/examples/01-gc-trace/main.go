package main

import (
	"fmt"
	"runtime"
	"time"
)

// allocateHeavy membuat beban di heap
func allocateHeavy() {
	for i := 0; i < 10000; i++ {
		// Alokasi slice kecil tiap iterasi
		_ = make([]byte, 10*1024)
	}
}

func main() {
	fmt.Println("Starting GC Intensive Application...")
	fmt.Println("GOMEMLIMIT set to standard")
	fmt.Println("GOGC set to 100 (default)")

	// Berikan waktu untuk user melihat status awal
	time.Sleep(1 * time.Second)

	// Lakukan alokasi berulang untuk memicu GC
	for i := 0; i < 100; i++ {
		allocateHeavy()
		time.Sleep(50 * time.Millisecond)
		
		if i%10 == 0 {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("Tick %d: HeapAlloc = %d MiB\n", i, m.HeapAlloc/1024/1024)
		}
	}

	fmt.Println("\nLab: Analisis Log GC")
	fmt.Println("Jalankan perintah berikut di terminal (mac/linux/windows):")
	fmt.Println("  $env:GODEBUG=\"gctrace=1\"; go run main.go")
	fmt.Println("\nArti Log GC:")
	fmt.Println("gc 1 @2.103s 5%: 0.051+1.5+0.003 ms clock, 0.41+1.5/1.2/0.17+0.024 ms cpu")
	fmt.Println("-> 2.103s: waktu sejak start")
	fmt.Println("-> 5%: penggunaan CPU oleh GC")
	fmt.Println("-> 0.051+1.5+0.003: durasi (STW Start + Concurrent Mark + STW End)")
}
