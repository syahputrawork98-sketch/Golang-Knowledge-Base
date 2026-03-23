package main

import "fmt"

func main() {
	// 1. Anonymous Struct untuk satu set data sementara
	point := struct {
		X, Y int
	}{
		X: 10,
		Y: 20,
	}

	fmt.Printf("Point: (%d, %d)\n", point.X, point.Y)

	// 2. Contoh penggunaan nyata: Table-driven Testing (Single run demo)
	testCases := []struct {
		input  string
		expect bool
	}{
		{"GO", true},
		{"JS", false},
	}

	for _, tc := range testCases {
		fmt.Printf("Checking %s: Expected %v\n", tc.input, tc.expect)
	}
}
