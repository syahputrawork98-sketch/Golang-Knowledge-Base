package main

import "fmt"

func main() {
	x := 10
	
	// Jebakan Defer: 
	// Nilai argumen 'x' di dalam fmt.Println dievaluasi SAAT pernyataan defer ditulis,
	// bukan saat fungsi main selesai.
	defer fmt.Printf("Nilai X di dalam Defer (Evaluasi Awal): %d\n", x)

	x = 20
	fmt.Printf("Nilai X di dalam Main (Setalah perubahan): %d\n", x)
}
