package main

import (
	"fmt"
	"time"
)

// SafeGo adalah pola pembungkus goroutine agar aman dari crash total.
func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[ALERTI] Goroutine terhenti karena error: %v\n", r)
			}
		}()
		fn()
	}()
}

func main() {
	fmt.Println("Menjalankan tugas di latar belakang...")

	SafeGo(func() {
		fmt.Println("Goroutine 1: Sedang bekerja...")
		panic("Kesalahan fatal di goroutine 1!")
	})

	SafeGo(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2: Saya tetap hidup karena G1 diisolasi.")
	})

	time.Sleep(2 * time.Second)
	fmt.Println("Program utama selesai.")
}
