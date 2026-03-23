package main

import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("--- RECOVERY AKTIF ---")
			fmt.Printf("Menangkap kegagalan fatal: %v\n", r)
			fmt.Println("Sistem dipulihkan, melanjutkan alur program...")
		}
	}()

	fmt.Println("Memulai operasi berisiko...")
	panic("DATABASE_CONNECTION_LOST") // Memicu panic manual
	fmt.Println("Ini tidak akan pernah dieksekusi")
}

func main() {
	riskyOperation()
	fmt.Println("Program utama tetap berjalan normal berkat 'recover'.")
}
