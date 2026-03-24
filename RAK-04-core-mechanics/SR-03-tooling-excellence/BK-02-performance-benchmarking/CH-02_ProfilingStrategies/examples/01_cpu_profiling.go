package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

// 01_cpu_profiling.go - Audit Jantung (Check-up)
// Jalankan lalu analisis dengan: go tool pprof cpu.prof

func heavyWork() {
	result := 0
	for i := 0; i < 100000000; i++ {
		result += i
	}
	fmt.Println(result)
}

func main() {
	// 1. Create Profile File
	f, _ := os.Create("cpu.prof")
	defer f.Close()

	// 2. Start Profiling
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println("Running heavy work...")
	heavyWork()
}
