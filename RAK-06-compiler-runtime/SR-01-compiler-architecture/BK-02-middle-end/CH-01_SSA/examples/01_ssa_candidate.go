package main

// 01_ssa_candidate.go
// Contoh kecil untuk melihat bagaimana compiler membangun SSA.
// Jalankan dengan environment seperti:
//   GOSSAFUNC=compute go build

func compute(a, b int) int {
	sum := a + b
	if sum > 10 {
		return sum * 2
	}
	return sum - 1
}

func main() {
	println(compute(4, 9))
}
