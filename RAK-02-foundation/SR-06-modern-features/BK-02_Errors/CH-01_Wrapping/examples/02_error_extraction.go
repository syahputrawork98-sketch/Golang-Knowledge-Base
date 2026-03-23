package main

import (
	"errors"
	"fmt"
)

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API Error [%d]: %s", e.Code, e.Message)
}

func RequestData() error {
	myErr := &ApiError{Code: 404, Message: "Not Found"}
	return fmt.Errorf("operation failed: %w", myErr)
}

func main() {
	err := RequestData()

	// Mencoba mengekstrak ke struct APIError
	var target *ApiError
	if errors.As(err, &target) {
		fmt.Println("Success extracting ApiError!")
		fmt.Printf("Status Code: %d\n", target.Code)
		fmt.Printf("Message    : %s\n", target.Message)
	} else {
		fmt.Println("Failed to extract ApiError.")
	}
}
