package main

import "fmt"

func main() {
	s := []int{1, 2}
	
	// SALAH: Memasukkan ke append tapi tidak menangkap hasilnya
	// LINT: 'append result not used' adalah INTENTIONAL di sini untuk demo kesalahan.
	append(s, 3) 
	fmt.Println("Hasil tanpa tangkap:", s) // Tetap [1, 2]

	// BENAR: Menangkap kembali hasil header baru
	s = append(s, 3)
	fmt.Println("Hasil dengan tangkap:", s) // [1, 2, 3]

	fmt.Println("\nIngat: append tidak mengubah slice asal di tempat (in-place) " +
		"jika harus realokasi, ia mengembalikan header baru!")
}
