package main

import "fmt"

// LOGIKA: Saya butuh sesuatu yang bisa 'Save'
type Saver interface {
	Save(data string)
}

func ProcessAndSave(s Saver, data string) {
	fmt.Println("Processing data...")
	s.Save(data)
}

// STORAGE (Producer): Saya hanya menyediakan fungsi Save
// Saya tidak perlu tahu tentang interface 'Saver' di atas.
type FileStorage struct{}
func (FileStorage) Save(d string) { fmt.Println("Saved to file:", d) }

func main() {
	fs := FileStorage{}
	
	// Interface terpenuhi secara otomatis (Implicit)
	// Meskipun FileStorage tidak pernah menyebutkan 'Saver'
	ProcessAndSave(fs, "Important Content")
	
	fmt.Println("\nKesimpulan: Definisikan interface di tempat ia digunakan (Consumer).")
}
