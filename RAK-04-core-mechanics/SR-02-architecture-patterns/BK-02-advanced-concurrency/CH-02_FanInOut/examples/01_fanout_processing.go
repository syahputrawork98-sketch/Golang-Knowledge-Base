package main

import (
	"fmt"
	"sync"
	"time"
)

// 01_fanout_processing.go - Paralelisasi Tugas (Fan-Out)
// Analogi: Tim penyisir lapangan yang berpencar mencari sampah.

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("Worker %d: Squaring %d\n", id, n)
			time.Sleep(500 * time.Millisecond)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	in := producer(1, 2, 3, 4, 5)

	// Fan-Out: Distribusikan pekerjaan ke 3 goroutine sq
	c1 := square(1, in)
	c2 := square(2, in)
	c3 := square(3, in)

	// Fan-In: Gabungkan hasil (Sederhana via WaitGroup)
	var wg sync.WaitGroup
	wg.Add(3)
	
	outputChan := make(chan int)
	go func() {
		for n := range c1 { outputChan <- n }
		wg.Done()
	}()
	go func() {
		for n := range c2 { outputChan <- n }
		wg.Done()
	}()
	go func() {
		for n := range c3 { outputChan <- n }
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	for res := range outputChan {
		fmt.Printf("Result: %d\n", res)
	}
}
