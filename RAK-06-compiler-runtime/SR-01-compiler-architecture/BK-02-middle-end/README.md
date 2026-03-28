# BK-02: Compiler Middle-End

Buku ini membahas lapisan tengah compiler Go, terutama representasi SSA yang dipakai untuk analisis dan optimisasi sebelum kode diturunkan menjadi instruksi yang lebih dekat ke mesin.

## Struktur

### [CH-01_SSA](./CH-01_SSA/)
SSA sebagai representasi perantara yang membantu compiler melakukan transformasi dan optimisasi.

## Boundary

- fokus pada representasi menengah dan transformasi internal compiler;
- membantu pembaca memahami jembatan antara AST dan backend;
- bukan tempat utama untuk parsing awal atau ABI mesin.

---
*Status: [x] Complete*
