package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// simulateWork simulates a CPU-intensive operation
func simulateWork(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	return res
}

// simulateAlloc simulates a heap-heavy operation
func simulateAlloc() [][]byte {
	var data [][]byte
	for i := 0; i < 100; i++ {
		// Alokasi 1MB per iterasi
		buf := make([]byte, 1024*1024)
		data = append(data, buf)
	}
	return data
}

func main() {
	// 1. CPU Profile
	cpuFile, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	log.Println("Starting CPU heavy work...")
	for i := 0; i < 50; i++ {
		simulateWork(1000000)
	}

	// 2. Heap Profile
	log.Println("Starting Heap heavy work...")
	data := simulateAlloc()
	_ = data // Keep in memory

	memFile, err := os.Create("heap.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer memFile.Close()

	runtime.GC() // Get clean stats
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal(err)
	}

	log.Println("Profiling complete. Run:")
	log.Println("  go tool pprof -http=:8080 cpu.pprof")
	log.Println("  go tool pprof -http=:8081 heap.pprof")
	
	time.Sleep(1 * time.Second)
}
