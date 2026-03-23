package main

import "fmt"

func Describe(i any) {
	// Variabel 'v' di bawah akan berubah tipe secara otomatis di setiap case
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer berlipat dua: %d\n", v*2)
	case string:
		fmt.Printf("String panjangnya: %d\n", len(v))
	case bool:
		fmt.Printf("Boolean terbalik: %t\n", !v)
	default:
		fmt.Printf("Tipe kustom: %T\n", v)
	}
}

func main() {
	Describe(21)
	Describe("Antigravity")
	Describe(true)
	Describe(3.14)
}
