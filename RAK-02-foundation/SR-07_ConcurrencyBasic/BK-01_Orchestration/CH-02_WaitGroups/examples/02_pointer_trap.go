package main

import (
	"fmt"
	"sync"
)

// SALAH: Menerima WaitGroup sebagai value (copy)
func badWorker(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Bad worker done (but main won't see it)")
}

// BENAR: Menerima WaitGroup sebagai pointer
func goodWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good worker done")
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("--- Testing Pointer (Correct) ---")
	wg.Add(1)
	go goodWorker(&wg)
	wg.Wait()
	fmt.Println("Main detected Good Worker finishing.")

	fmt.Println("\n--- Testing Value (Incorrect - This would Deadlock) ---")
	fmt.Println("Note: Baris di bawah ini sengaja tidak saya jalankan karena " +
		"akan menyebabkan fatal error: all goroutines are asleep (deadlock).")
	
	// wg.Add(1)
	// go badWorker(wg) // wg ini dicopy, Done() tidak mempengaruhi wg asli di main
	// wg.Wait()
}
