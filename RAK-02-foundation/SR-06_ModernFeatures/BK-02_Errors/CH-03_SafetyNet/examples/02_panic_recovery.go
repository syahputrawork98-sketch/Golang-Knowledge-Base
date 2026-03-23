package main

import "fmt"

func SafeExecution() {
	// Mendaftarkan 'Safety Net'
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n[RECOVERY] Sistem berhasil menangkap ledakan (panic)!")
			fmt.Printf("[REASON] %v\n", r)
			fmt.Println("[STATUS] Melanjutkan eksekusi dengan aman.")
		}
	}()

	fmt.Println("Memulai operasi berisiko tinggi...")
	
	// Simulasi kecelakaan runtime (Nil Pointer)
	var ptr *int
	fmt.Println("Nilai:", *ptr) // Ini akan memicu PANIC

	fmt.Println("Baris ini tidak akan pernah dieksekusi.")
}

func main() {
	fmt.Println("--- START PROGRAM ---")
	SafeExecution()
	fmt.Println("--- END PROGRAM (Successfully Survived) ---")
}
