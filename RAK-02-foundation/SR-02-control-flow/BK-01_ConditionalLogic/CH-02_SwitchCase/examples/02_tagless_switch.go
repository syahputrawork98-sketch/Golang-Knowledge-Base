package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Tagless Switch (Bertindak sebagai rangkaian if-else)
	// Sangat idiomatik untuk mengecek logika yang tidak berhubungan dengan satu variabel saja.
	hour := time.Now().Hour()

	fmt.Printf("Pesan Waktu (%d:00): ", hour)
	switch {
	case hour < 12:
		fmt.Println("Selamat Pagi!")
	case hour < 17:
		fmt.Println("Selamat Siang!")
	default:
		fmt.Println("Selamat Malam!")
	}

	// 2. Multi-value Case
	day := "Sabtu"
	switch day {
	case "Sabtu", "Minggu":
		fmt.Println("Status: Akhir Pekan (Libur)")
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Status: Hari Kerja")
	default:
		fmt.Println("Status: Hari tidak valid")
	}
}
