package main

import "fmt"

func main() {
	var val interface{} = "Antigravity"

	// 1. Assertion Berbahaya (Bisa Panic!)
	// s := val.(int) 

	// 2. Assertion Aman (Comma-ok)
	if s, ok := val.(string); ok {
		fmt.Println("Value adalah string:", s)
	} else {
		fmt.Println("Bukan string!")
	}

	// 3. Mencoba ke tipe lain
	if f, ok := val.(float64); ok {
		fmt.Println("Value adalah float:", f)
	} else {
		fmt.Println("Gagal mengonversi ke float")
	}
}
