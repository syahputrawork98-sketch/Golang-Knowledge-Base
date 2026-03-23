package main

import "fmt"

type NetworkError struct {
	Temporary bool
}

func (e *NetworkError) Error() string {
	return "Network problem"
}

// Method tambahan untuk mendeteksi perilaku
func (e *NetworkError) IsRetryable() bool {
	return e.Temporary
}

func main() {
	err := &NetworkError{Temporary: true}

	// Alih-alih cek tipe, kita cek perilakunya (Behavioral checking)
	if netErr, ok := interface{}(err).(interface{ IsRetryable() bool }); ok {
		if netErr.IsRetryable() {
			fmt.Println("System Action: Retrying connection...")
		} else {
			fmt.Println("System Action: Aborting operation.")
		}
	}
}
