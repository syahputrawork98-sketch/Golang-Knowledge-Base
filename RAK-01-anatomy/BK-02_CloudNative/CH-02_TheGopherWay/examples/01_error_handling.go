package main

import (
	"errors"
	"fmt"
)

// Fungsi idiomatik Go: mengembalikan (hasil, error)
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("--- The Gopher Way: Explicit Error Check ---")

	// Panggilan 1: Berhasil
	val, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Error occurred: %s\n", err)
	} else {
		fmt.Printf("Result: %v\n", val)
	}

	// Panggilan 2: Gagal
	val2, err2 := Divide(10, 0)
	if err2 != nil {
		fmt.Printf("Explicit Error: %s\n", err2)
	} else {
		fmt.Printf("Result: %v\n", val2)
	}

	fmt.Println("Program selesai dengan linear tanpa exception jump.")
}
