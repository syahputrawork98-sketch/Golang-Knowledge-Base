package main

import "fmt"

type Storage interface {
	Save(data string) error
}

type Processor struct {
	repo Storage
}

func (p Processor) Execute(data string) {
	err := p.repo.Save(data)
	if err != nil {
		fmt.Println("Error detected in test/prod:", err)
		return
	}
	fmt.Println("Execution Successful!")
}

// --- MOCK OBJECT UNTUK TESTING ---
type MockStorage struct {
	ShouldFail bool
}

func (m MockStorage) Save(data string) error {
	if m.ShouldFail {
		return fmt.Errorf("simulated storage failure")
	}
	fmt.Println("[MOCK] Saving data to memory...")
	return nil
}

func main() {
	fmt.Println("--- Scenario 1: Test Success ---")
	testSuccess := Processor{repo: MockStorage{ShouldFail: false}}
	testSuccess.Execute("Data OK")

	fmt.Println("\n--- Scenario 2: Test Failure ---")
	testFail := Processor{repo: MockStorage{ShouldFail: true}}
	testFail.Execute("Bad Data")
	
	fmt.Println("\nKesimpulan: Interface memungkinkan kita mensimulasikan 'Error' tanpa merusak database asli.")
}
