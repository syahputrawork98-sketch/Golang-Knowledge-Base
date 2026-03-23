package main

import "fmt"

func main() {
	var val any = "Mastering Go"

	// 1. Assertion Berbahaya (Bisa Panic!)
	// n := val.(int) // panic: interface conversion: interface {} is string, not int

	// 2. Assertion Aman (Pola Comma-OK)
	if s, ok := val.(string); ok {
		fmt.Printf("Sukses: Value adalah string '%s'\n", s)
	}

	if n, ok := val.(int); !ok {
		fmt.Println("Gagal: Value bukan integer, tapi program tetap berjalan (tidak panic)")
		fmt.Printf("Zero value 'n': %d\n", n)
	}
}
