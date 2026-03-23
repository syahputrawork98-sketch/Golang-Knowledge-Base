package main

import "fmt"

func main() {
	// Map yang nil (hanya deklarasi)
	var nilMap map[string]int
	
	// Membaca nil map: AMAN (hasilnya zero value)
	fmt.Println("Baca nilMap:", nilMap["any"])

	// Menulis ke nil map: PANIC! (Uncomment baris bawah untuk tes)
	// nilMap["key"] = 1 

	// Map kosong (sudah diinisialisasi)
	emptyMap := make(map[string]int)
	emptyMap["key"] = 1 // AMAN
	
	fmt.Println("Beres. Gunakan make() sebelum mengisi map!")
}
