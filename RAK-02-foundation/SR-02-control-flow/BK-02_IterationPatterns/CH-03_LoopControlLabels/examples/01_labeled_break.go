package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	target := 5

	fmt.Println("Mencari angka", target, "di dalam matriks...")

// Label 'OuterLoop' menempel pada perulangan terluar.
OuterLoop:
	for r, row := range matrix {
		for c, val := range row {
			if val == target {
				fmt.Printf("Ditemukan! Baris: %d, Kolom: %d\n", r, c)
				break OuterLoop // Langsung keluar dari seluruh perulangan bertingkat.
			}
			fmt.Printf("Mengecek [%d][%d]...\n", r, c)
		}
	}

	fmt.Println("Pencarian selesai.")
}
