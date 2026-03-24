package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// 01_inspect_ast.go - Membedah Pohon Sintaksis
// Analogi: Menggunakan X-Ray untuk melihat susunan tulang dari sebuah kalimat.

func main() {
	fset := token.NewFileSet()

	// Kode sumber Go yang akan kita bedah
	src := `package main
func hello() {
	println("Hello Gopher!")
}`

	// 1. Parsing kode menjadi AST
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	// 2. Traversal AST (Melewati setiap node di pohon)
	ast.Inspect(f, func(n ast.Node) bool {
		// Jika node adalah deklarasi fungsi
		if fn, ok := n.(*ast.FuncDecl); ok {
			fmt.Printf("Found function: %s (at line %d)\n", fn.Name.Name, fset.Position(fn.Pos()).Line)
		}
		
		// Jika node adalah literal string
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			fmt.Printf("Found string constant: %s\n", lit.Value)
		}
		
		return true
	})
}
