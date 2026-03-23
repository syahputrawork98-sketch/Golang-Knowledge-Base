package main

import (
	"fmt"
	"strings"
)

// ValidationError menampung banyak kesalahan field dalam satu objek
type ValidationError struct {
	Fields map[string]string
}

func (e *ValidationError) Error() string {
	var msgs []string
	for field, msg := range e.Fields {
		msgs = append(msgs, fmt.Sprintf("%s: %s", field, msg))
	}
	return "Validation failed: " + strings.Join(msgs, ", ")
}

func ValidateUser(name string, age int) error {
	fields := make(map[string]string)
	if name == "" {
		fields["name"] = "cannot be empty"
	}
	if age < 18 {
		fields["age"] = "must be at least 18"
	}

	if len(fields) > 0 {
		return &ValidationError{Fields: fields}
	}
	return nil
}

func main() {
	err := ValidateUser("", 10)
	if err != nil {
		fmt.Println("User seeing this:", err)
		
		// Teknik Senior: Type Assertion untuk akses metadata
		if vErr, ok := err.(*ValidationError); ok {
			fmt.Println("\nDeveloper logs:")
			for f := range vErr.Fields {
				fmt.Println("- Broken field:", f)
			}
		}
	}
}
