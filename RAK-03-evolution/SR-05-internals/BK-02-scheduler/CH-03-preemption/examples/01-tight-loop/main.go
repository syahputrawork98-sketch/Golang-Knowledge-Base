package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 1. Batasi Processor menjadi 1 untuk simulasi penguncian (hogging)
	runtime.GOMAXPROCS(1)
	fmt.Println("GOMAXPROCS set to 1")

	// 2. Luncurkan goroutine "pelayan" yang ramah
	go func() {
		for {
			fmt.Println("Hello from friendly goroutine!")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 3. Luncurkan goroutine "rakus" (tight loop)
	// Di Go versi lama (< 1.14), goroutine ini akan mengunci P selamanya.
	// Di Go versi baru (1.14+), dia akan diputus paksa (preempted).
	go func() {
		fmt.Println("HOGGING START: Greedy goroutine is running a tight loop...")
		i := 0
		for {
			i++
			// Loop ini tidak memanggil fungsi apapun (tidak ada safe point).
		}
	}()

	// Tunggu 5 detik untuk melihat apakah goroutine ramah tetap jalan.
	time.Sleep(5 * time.Second)
	fmt.Println("\nLab Success: Friendly goroutine was NOT blocked!")
	fmt.Println("Conclusion: Asynchronous Preemption (Go 1.14+) is working.")
}
