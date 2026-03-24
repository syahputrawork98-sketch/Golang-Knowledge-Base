package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 01_simple_api_server.go - Membangun JSON API Sederhana
// Analogi: Koki Restoran yang menyajikan menu dalam format standar.

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Menentukan header content-type
	w.Header().Set("Content-Type", "application/json")
	
	resp := Response{
		Status:  "Success",
		Message: "Hello from Gopher Standard Library!",
	}
	
	// Encode struct langsung ke writer
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	
	fmt.Println("Server starts at :8080...")
	// http.ListenAndServe(":8080", nil) // Jalankan manual
}
