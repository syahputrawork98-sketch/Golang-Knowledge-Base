package main

import "fmt"

// Fungsi ini menerima array sebagai VALUE (salinan)
func modifyArray(a [3]int) {
	a[0] = 999
	fmt.Println("Di dalam fungsi:", a)
}

func main() {
	original := [3]int{1, 2, 3}
	
	fmt.Println("Sebelum fungsi:", original)
	
	modifyArray(original)
	
	// 'original' tidak berubah karena Go melakukan full copy ke parameter fungsi
	fmt.Println("Setelah fungsi:", original)
	
	fmt.Println("\nKesimpulan: Di Go, Array adalah Value Type. Hati-hati menyalin array besar!")
}
