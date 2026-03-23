package main

import "fmt"

// Fungsi untuk mengekstrak key dari map apa saja
// K harus comparable karena map key butuh perbandingan ==
func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	// Map dengan key string
	m1 := map[string]int{"a": 1, "b": 2}
	fmt.Println("Keys m1:", GetKeys(m1))

	// Map dengan key int
	m2 := map[int]string{1: "Admin", 2: "User"}
	fmt.Println("Keys m2:", GetKeys(m2))
}
