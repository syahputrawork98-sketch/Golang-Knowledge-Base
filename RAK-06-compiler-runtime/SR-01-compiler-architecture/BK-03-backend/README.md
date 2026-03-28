# BK-03: Compiler Backend

Buku ini membahas tahap akhir compiler Go: bagaimana representasi internal diterjemahkan menjadi assembly, object file, dan akhirnya executable binary.

## Struktur

### [CH-01_CodeGen](./CH-01_CodeGen/)
Code generation, assembly output, dan peran ABI dalam tahap backend compiler.

## Boundary

- fokus pada penurunan representasi compiler ke bentuk yang bisa dieksekusi mesin;
- membantu pembaca memahami hubungan antara SSA, assembly, dan binary final;
- bukan tempat utama untuk scheduler runtime atau allocator memori.

---
*Status: [x] Complete*
