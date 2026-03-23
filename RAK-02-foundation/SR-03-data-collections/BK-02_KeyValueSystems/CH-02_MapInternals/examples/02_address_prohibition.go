package main

import "fmt"

func main() {
	m := map[string]int{"age": 25}

	// ERROR: cannot take the address of m["age"]
	// Karena map bisa realokasi (growing) kapan saja, 
	// menunjuk ke alamat internal map sangatlah berbahaya.
	
	// ptr := &m["age"] 
	
	fmt.Println("Nilai m['age']:", m["age"])
	fmt.Println("Catatan: Anda tidak bisa mengambil pointer &m[key] di Go!")
}
