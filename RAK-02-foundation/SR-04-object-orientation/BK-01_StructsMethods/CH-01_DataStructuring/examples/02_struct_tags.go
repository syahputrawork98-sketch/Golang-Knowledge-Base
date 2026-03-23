package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
	Stock int     `json:"-"` // Field ini akan diabaikan oleh JSON
}

func main() {
	p := Product{
		Name:  "Mechanical Keyboard",
		Price: 150.50,
		Stock: 50,
	}

	// Marshalling: Struct -> JSON
	jsonData, _ := json.Marshal(p)
	fmt.Println("JSON Result:", string(jsonData))

	// Unmarshalling: JSON -> Struct
	jsonStr := `{"product_name": "Gaming Mouse", "price": 80.00}`
	var p2 Product
	json.Unmarshal([]byte(jsonStr), &p2)

	fmt.Printf("Parsed Product: %+v\n", p2)
	fmt.Println("Stock is always 0 (ignored):", p2.Stock)
}
