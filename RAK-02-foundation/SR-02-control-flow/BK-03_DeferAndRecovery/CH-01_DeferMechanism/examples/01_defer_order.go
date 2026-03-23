package main

import "fmt"

func main() {
	fmt.Println("Memulai pendaftaran Defer...")

	// 1. Sifat LIFO (Last-In-First-Out)
	// Yang terakhir didefer akan dijalankan pertama kali.
	defer fmt.Println("Eksekusi 1 (Didaftar paling awal)")
	defer fmt.Println("Eksekusi 2")
	defer fmt.Println("Eksekusi 3 (Didaftar paling akhir)")

	fmt.Println("Proses sedang berjalan...")
	fmt.Println("Fungsi akan segera selesai. Perhatikan urutan output di bawah ini:")
}
