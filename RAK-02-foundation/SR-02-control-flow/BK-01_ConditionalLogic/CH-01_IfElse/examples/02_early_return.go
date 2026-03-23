package main

import (
	"fmt"
)

// Pola "Happy Path on the Left" adalah standar Senior Go.
// Kita menangani kondisi error/pengecualian secepat mungkin (Early Return).
func validateUser(username string, age int) error {
	if username == "" {
		return fmt.Errorf("username tidak boleh kosong")
	}

	if age < 18 {
		return fmt.Errorf("user harus berusia minimal 18 tahun")
	}

	// Happy Path: Logika utama tetap berada di identasi terluar (sisi kiri).
	// Ini jauh lebih mudah dibaca daripada nested if-else yang dalam.
	fmt.Printf("User %s valid dan siap diproses.\n", username)
	return nil
}

func main() {
	err := validateUser("antigravity", 25)
	if err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Println("Proses selesai.")
}
