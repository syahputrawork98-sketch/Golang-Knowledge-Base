package main

import (
	"crypto/sha256"
	"fmt"
)

// 01_password_hashing.go - Integritas Data (Hashing)
// Analogi: Mesin Penghancur Kertas yang menghasilkan pola unik.

func hashPassword(password string, salt string) string {
	// Menggabungkan password dan salt untuk keamanan ekstra
	h := sha256.New()
	h.Write([]byte(password + salt))
	
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	pass := "gopher_kuat_123"
	salt := "random_string_xyz"

	hashed := hashPassword(pass, salt)
	fmt.Println("Original Password:", pass)
	fmt.Println("SHA256 Hash (with Salt):", hashed)
}
