package main

import "fmt"

func Inspect(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Ini adalah Integer: %d\n", v)
	case string:
		fmt.Printf("Ini adalah String: %s\n", v)
	case bool:
		fmt.Printf("Ini adalah Boolean: %v\n", v)
	default:
		fmt.Printf("Tipe tidak dikenal: %T\n", i)
	}
}

func main() {
	Inspect(100)
	Inspect("Golang")
	Inspect(true)
	Inspect(3.14)
}
