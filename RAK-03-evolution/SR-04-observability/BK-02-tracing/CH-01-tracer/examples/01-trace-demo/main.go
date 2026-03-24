package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Simulasi kerja: sleep (blocking) then calculate (running)
	time.Sleep(10 * time.Millisecond)
	
	res := 0
	for i := 0; i < 1000000; i++ {
		res += i
	}
}

func main() {
	// 1. Inisialisasi Trace File
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer f.Close()

	// 2. Start Tracing
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	log.Println("Tracing started. Simulating concurrent orchestration...")
	
	// 3. Jalankan pekerjaan konkuren
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU() * 2
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	log.Println("Tracing complete. Run:")
	log.Println("  go tool trace trace.out")
}
