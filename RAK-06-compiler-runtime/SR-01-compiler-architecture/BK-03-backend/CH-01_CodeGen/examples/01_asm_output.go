package main

// 01_asm_output.go - Melihat Hasil Kompilasi ke Assembly
// Instruksi: Jalankan 'go tool compile -S 01_asm_output.go'

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(10, 20)
	println(result)
}
