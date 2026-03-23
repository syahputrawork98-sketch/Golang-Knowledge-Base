package main

import "fmt"

func identifyType(i interface{}) {
	// Type Switch menggunakan sintaks khusus: i.(type)
	// Hanya bisa digunakan di dalam konstruksi switch.
	switch v := i.(type) {
	case int:
		fmt.Printf("Ini adalah Integer: %d\n", v)
	case string:
		fmt.Printf("Ini adalah String: %q\n", v)
	case bool:
		fmt.Printf("Ini adalah Boolean: %t\n", v)
	default:
		fmt.Printf("Tipe tidak dikenal oleh sistem: %T\n", v)
	}
}

func main() {
	identifyType(100)
	identifyType("Halo Go")
	identifyType(true)
	identifyType(3.14) // float64 akan masuk ke default
}
