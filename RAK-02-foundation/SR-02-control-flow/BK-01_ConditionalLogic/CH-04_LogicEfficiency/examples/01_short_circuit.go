package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Memulai pengujian Short-Circuiting (&&)...")

	// Kasus 1: Kondisi pertama False.
	// Fungsi heavyTask() TIDAK AKAN dijalankan sama sekali.
	if false && heavyTask() {
		fmt.Println("Ini tidak akan tercetak")
	} else {
		fmt.Println("Kondisi pertama False, pengecekan langsung berhenti (Efisien).")
	}

	fmt.Println("\nMemulai pengujian Short-Circuiting (||)...")

	// Kasus 2: Kondisi pertama True.
	// Fungsi heavyTask() juga TIDAK AKAN dijalankan karena hasil sudah pasti True.
	if true || heavyTask() {
		fmt.Println("Kondisi pertama True, pengecekan langsung berhenti (Efisien).")
	}
}

func heavyTask() bool {
	fmt.Println(">>> [LOG] Sedang menjalankan tugas berat selama 2 detik...")
	time.Sleep(2 * time.Second)
	return true
}
