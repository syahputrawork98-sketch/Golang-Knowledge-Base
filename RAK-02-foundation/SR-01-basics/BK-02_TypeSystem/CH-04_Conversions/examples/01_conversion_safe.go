package main

import "fmt"

func main() {
	fmt.Println("--- Explicit Type Conversion in Go ---")

	var x int = 42
	var y float64 = float64(x) // Explicit conversion
	var z uint = uint(y)

	fmt.Printf("Initial int: %v, Converted float: %v, Back to uint: %v\n", x, y, z)

	// Demonstrasi Pemotongan Data (Truncation)
	pi := 3.14159
	integerPi := int(pi)
	fmt.Printf("Pi: %v -> int(Pi): %v (Data di belakang koma hilang!)\n", pi, integerPi)

	// Demonstrasi Bahaya Overflow
	var largeInt int64 = 500
	var smallInt int8 = int8(largeInt)
	fmt.Printf("int64(500) -> int8: %v (Overflow terjadi!)\n", smallInt)
}
