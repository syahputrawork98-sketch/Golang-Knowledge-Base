package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 02_http_client_post.go - Mengirim Pesanan ke Server
// Analogi: Pelayan yang mengantarkan formulir pesanan ke dapur.

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	
	// 1. Siapkan data yang akan dikirim
	data := map[string]string{
		"title":  "Gopher Standard Library",
		"body":   "Learning HTTP Client at Platinum Gold Standard",
		"userId": "1",
	}
	jsonData, _ := json.Marshal(data)

	// 2. Buat Request dengan Timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request Error:", err)
		return
	}
	defer resp.Body.Close()

	// 3. Baca Response
	fmt.Println("Status Code:", resp.Status)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}
