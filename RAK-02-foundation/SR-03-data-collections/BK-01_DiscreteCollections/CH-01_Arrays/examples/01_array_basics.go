package main

import "fmt"

func main() {
	// 1. Deklarasi dengan ukuran eksplisit
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20

	// 2. Deklarasi dengan inisialisasi (Array Literal)
	colors := [3]string{"Red", "Green", "Blue"}

	// 3. Menggunakan ellipsis (...) agar compiler menghitung ukurannya
	prices := [...]float64{10.5, 20.0, 15.75, 40.2}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Colors:", colors)
	fmt.Printf("Prices (Size: %d): %v\n", len(prices), prices)

	// 4. Perulangan standar
	fmt.Println("\nIterasi Array:")
	for i, v := range colors {
		fmt.Printf("Index %d has color %s\n", i, v)
	}
}
