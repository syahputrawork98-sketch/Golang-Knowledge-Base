package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 1. Tunjukkan jumlah P (Processor) default
	numP := runtime.GOMAXPROCS(0)
	fmt.Printf("Default GOMAXPROCS (P): %d\n", numP)

	// 2. Simulasi beban kerja berat untuk melihat G-M-P beraksi
	var wg sync.WaitGroup
	numG := 100 // 100 Goroutines

	fmt.Printf("Launching %d Goroutines across %d Processors...\n", numG, numP)

	for i := 0; i < numG; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Simulasi komputasi intensif
			count := 0
			for j := 0; j < 1000000; j++ {
				count += j
			}
			
			if id%10 == 0 {
				fmt.Printf("[G-%d] Running on OS Thread\n", id)
			}
		}(i)
	}

	// 3. Monitor jumlah thread OS (M) yang aktif
	go func() {
		for {
			numM, _ := runtime.ThreadCreateProfile(nil)
			fmt.Printf("--> Monitoring: Active M (OS Threads): %d\n", numM)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Simulation Complete.")
	fmt.Println("\nLab Guide:")
	fmt.Println("1. Perhatikan bahwa M (OS Threads) biasanya > P (Processor) ")
	fmt.Println("   karena ada thread untuk GC, Syscall, dll.")
	fmt.Println("2. G ditampung di Local Run Queue (LRQ) milik masing-masing P.")
}
