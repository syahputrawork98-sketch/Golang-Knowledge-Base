package main

import "fmt"

// 01_data_pipeline.go - Aliran Data Terpadu (Stage-by-Stage)
// Analogi: Baris perakitan mobil.

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func multiplier(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 10
		}
		close(out)
	}()
	return out
}

func main() {
	// Setup Pipeline: Gen -> Filter -> Multiply
	source := generator(1, 2, 3, 4, 5, 6)
	filtered := filterEven(source)
	final := multiplier(filtered)

	for res := range final {
		fmt.Printf("Pipeline Result: %d\n", res)
	}
}
