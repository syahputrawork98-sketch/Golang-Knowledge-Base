package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
	IsActive bool
}

func main() {
	// 1. Inisialisasi dengan Struct Literal (Rekomendasi)
	u1 := User{
		ID:       1,
		Username: "antigravity",
		Email:    "ai@google.com",
		IsActive: true,
	}

	// 2. Short literal (Hati-hati: urutan field harus tepat)
	u2 := User{2, "godev", "dev@go.dev", false}

	// 3. Pointer ke Struct
	u3 := &User{ID: 3, Username: "pointer_user"}

	fmt.Printf("User 1: %+v\n", u1)
	fmt.Printf("User 2 Name: %s\n", u2.Username)
	fmt.Printf("User 3 Type: %T\n", u3)

	// 4. Modifikasi melalui pointer (Otomatis didereferensi oleh Go)
	u3.IsActive = true
	fmt.Printf("User 3 Active: %v\n", u3.IsActive)
}
