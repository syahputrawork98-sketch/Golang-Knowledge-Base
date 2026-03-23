package reverse

// Reverse membalik urutan string berbasis rune (UTF-8 safe)
func Reverse(s string) (string, error) {
	// Implementasi yang mungkin mengandung bug jika inputnya byte mentah / non-UTF8
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
