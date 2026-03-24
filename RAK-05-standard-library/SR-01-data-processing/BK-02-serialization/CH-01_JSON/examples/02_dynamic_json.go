package main

import (
	"encoding/json"
	"fmt"
)

// 02_dynamic_json.go - Menangani JSON Dinamis
// Analogi: Membuka kotak misteri tanpa tahu isinya.

func main() {
	jsonRaw := `{
		"event": "login",
		"timestamp": 1679645600,
		"meta": {
			"ip": "192.168.1.1",
			"browser": "Chrome"
		}
	}`

	// Menggunakan map[string]interface{} untuk struktur yang tidak pasti
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonRaw), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Event Name:", data["event"])
	
	// Akses nested data dengan type assertion
	meta := data["meta"].(map[string]interface{})
	fmt.Println("IP Address:", meta["ip"])
}
