package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Map adalah tipe data 'any' (interface{})
	var sm sync.Map

	// Store (Simpan)
	sm.Store("user_1", "Antigravity")
	sm.Store("user_2", "GoExpert")

	// Load (Baca)
	val, ok := sm.Load("user_1")
	if ok {
		fmt.Println("Ditemukan:", val)
	}

	// Range (Iterasi)
	sm.Range(func(key, value any) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // lanjut iterasi
	})
	
	fmt.Println("sync.Map cocok untuk data yang jarang berubah (stable keys).")
}
