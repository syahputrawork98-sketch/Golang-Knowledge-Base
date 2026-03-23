package api

import (
	"encoding/json"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FetchUserResponse simulasi output API yang kompleks
func FetchUserResponse() ([]byte, error) {
	u := User{ID: 1, Name: "Gopher", Email: "gopher@golang.org"}
	return json.MarshalIndent(u, "", "  ")
}
