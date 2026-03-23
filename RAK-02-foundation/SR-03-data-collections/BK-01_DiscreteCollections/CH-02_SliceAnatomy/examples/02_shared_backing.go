package main

import "fmt"

func main() {
	// Menyiapkan backing array melalui slice awal
	names := []string{"Alpha", "Bravo", "Charlie", "Delta"}
	
	// Membuat sub-slice (Hanya menunjuk ke bagian dari array names)
	sub := names[1:3] // Bravo, Charlie (Index 1 hingga sebelum 3)
	
	fmt.Println("Original:", names)
	fmt.Println("Sub-slice:", sub)

	// Mengubah isi sub-slice
	sub[0] = "ZULU"
	
	fmt.Println("\nSetelah sub[0] diubah:")
	fmt.Println("Sub-slice:", sub)
	fmt.Println("Original (UTAMA):", names)

	fmt.Println("\nKesimpulan: Sub-slice berbagi Backing Array yang sama. " + 
		"Hati-hati, perubahan di satu tempat berdapak ke tempat lain!")
}
