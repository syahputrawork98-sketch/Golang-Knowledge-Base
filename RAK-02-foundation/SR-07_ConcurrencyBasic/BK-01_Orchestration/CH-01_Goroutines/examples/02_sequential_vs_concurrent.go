package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d started...\n", id)
	time.Sleep(1 * time.Second)
}

func main() {
	start := time.Now()

	fmt.Println("--- RUNNING SEQUENTIAL ---")
	for i := 1; i <= 3; i++ {
		task(i)
	}
	fmt.Printf("Sequential took: %v\n\n", time.Since(start))

	start = time.Now()
	fmt.Println("--- RUNNING CONCURRENT ---")
	for i := 1; i <= 3; i++ {
		go task(i)
	}
	
	// Kita butuh penahan agar main tidak langsung selesai
	// (Penyelesaian yang benar ada di bab WaitGroup)
	time.Sleep(1200 * time.Millisecond) 
	
	fmt.Printf("Concurrent took: %v\n", time.Since(start))
}
