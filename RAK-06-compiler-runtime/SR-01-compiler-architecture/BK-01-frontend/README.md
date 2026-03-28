# BK-01: Compiler Frontend

Buku ini membahas tahap paling awal dalam pipeline compiler Go, yaitu bagaimana source code dibaca, diparse, lalu diubah menjadi struktur yang bisa dianalisis lebih lanjut.

## Struktur

### [CH-01_ParserAST](./CH-01_ParserAST/)
Parser dan AST sebagai pintu masuk compiler untuk memahami bentuk program Go.

## Boundary

- fokus pada pembacaan dan pembentukan struktur program di awal kompilasi;
- membantu pembaca memahami transisi dari teks mentah ke representasi sintaksis;
- bukan tempat utama untuk optimisasi SSA atau code generation.

---
*Status: [x] Complete*
