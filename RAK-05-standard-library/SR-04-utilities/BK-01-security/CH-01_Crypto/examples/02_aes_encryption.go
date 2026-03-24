package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// 02_aes_encryption.go - Enkripsi Simetris (Brankas Kunci)
// Analogi: Memasukkan emas ke brankas dan menguncinya dengan kunci rahasia.

func main() {
	plaintext := []byte("Ini adalah pesan sangat rahasia dari Gopher!")
	key := []byte("kunci-rahasia-32-byte-level-gold") // AES-256 (32 bytes)

	// 1. Siapkan Block Cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Cipher Error:", err)
		return
	}

	// 2. GCM (Galois/Counter Mode) untuk keamanan modern
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	// 3. Seal (Enkripsi)
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("Ciphertext (Hex): %x\n", ciphertext)

	// 4. Open (Dekripsi)
	decrypted, _ := gcm.Open(nil, nonce, ciphertext[gcm.NonceSize():], nil)
	fmt.Printf("Decrypted Message: %s\n", string(decrypted))
}
