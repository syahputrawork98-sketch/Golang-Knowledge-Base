package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Aktifkan pprof via HTTP
	"runtime"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

// heavilyContendedFunction simulates high mutex contention
func heavilyContendedFunction() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		// Simulasi "logika berat" di dalam lock
		counter++
		time.Sleep(1 * time.Microsecond)
		mu.Unlock()
	}
}

func main() {
	// 1. Jalankan profiling contention (wajib disetel rate-nya)
	runtime.SetMutexProfileFraction(1) // Catat 100% kejadian contention
	
	// Jalankan pprof server
	go func() {
		log.Println("Pprof server started at http://localhost:6060/debug/pprof")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Simulating high mutex contention...")
	
	const numGoroutines = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			heavilyContendedFunction()
		}()
	}

	fmt.Println("Contention simulation running...")
	fmt.Println("Run:")
	fmt.Println("  go tool pprof http://localhost:6060/debug/pprof/mutex")
	
	wg.Wait()
	fmt.Println("Simulation complete. Counter:", counter)
}
