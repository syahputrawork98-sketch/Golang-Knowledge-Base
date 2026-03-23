package main

import "fmt"

func main() {
	// Map yang bisa menampung apa saja
	config := make(map[string]any)

	config["port"] = 8080
	config["secure"] = true
	config["name"] = "AppServer"

	for key, val := range config {
		fmt.Printf("Key: %s, Value: %v (Type: %T)\n", key, val, val)
	}

	fmt.Println("\nKesimpulan: 'any' sangat fleksibel tapi mengabaikan Type Safety statis.")
}
