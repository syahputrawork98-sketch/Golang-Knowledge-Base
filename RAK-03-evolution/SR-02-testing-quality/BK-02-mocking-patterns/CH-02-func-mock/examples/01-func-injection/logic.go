package funcinjection

import "fmt"

// GreeterLogic menerima fungsi sebagai dependensi
func GreeterLogic(id int, fetchName func(int) (string, error)) (string, error) {
	name, err := fetchName(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}
