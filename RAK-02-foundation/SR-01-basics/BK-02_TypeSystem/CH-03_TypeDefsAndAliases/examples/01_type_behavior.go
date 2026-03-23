package main

import "fmt"

// 1. Type Definition (Berbeda secara sistem)
type Kilogram float64

// 2. Type Alias (Sama secara sistem)
type Berat = float64

// Menambahkan method pada Type Definition
func (kg Kilogram) ToGram() float64 {
	return float64(kg * 1000)
}

func main() {
	fmt.Println("--- Type Definition vs Alias ---")

	var beratAsli float64 = 10.5
	var k Kilogram = 5
	var b Berat = 10.5

	// ALIAS: Bisa langsung dijumlahkan dengan tipe aslinya
	totalAlias := beratAsli + b
	fmt.Printf("Alias Total: %v (Direct math possible)\n", totalAlias)

	// DEFINITION: Tidak bisa dijumlahkan langsung (Error!)
	// totalDef := beratAsli + k // UNCOMMENT INI AKAN ERROR

	// Harus konversi eksplisit
	totalDef := Kilogram(beratAsli) + k
	fmt.Printf("Definition Total: %v (Needs explicit conversion)\n", totalDef)
	fmt.Printf("5 Kg to Gram: %v\n", k.ToGram())
}
