package main

import (
	"encoding/json"
	"fmt"
)

// 01_basic_marshal.go - Konversi Struct ke JSON
// Analogi: Membuat Paspor Internasional dari KTP Lokal.

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"contact_email"`
	Secret   string `json:"-"` // Tidak akan diikutkan dalam JSON
}

func main() {
	u := User{
		ID:       101,
		Username: "artdarkman",
		Email:    "art@gopher.com",
		Secret:   "password123",
	}

	// 1. Marshal: Struct -> JSON
	jsonData, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("JSON Result:\n", string(jsonData))

	// 2. Unmarshal: JSON -> Struct
	jsonInput := `{"id": 102, "username": "syahputra", "contact_email": "syah@work.id"}`
	var u2 User
	json.Unmarshal([]byte(jsonInput), &u2)
	fmt.Printf("\nUnmarshaled Struct: %+v\n", u2)
}
