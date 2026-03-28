# SR-01: Compiler Architecture

Sub-rak ini membahas pipeline compiler Go dari teks mentah hingga instruksi mesin.

## Struktur

### [BK-01-frontend](./BK-01-frontend/)
Parser, AST, dan tahap awal memahami struktur program.

### [BK-02-middle-end](./BK-02-middle-end/)
SSA dan transformasi menengah sebelum kode diturunkan lebih jauh.

### [BK-03-backend](./BK-03-backend/)
Code generation dan tahap akhir menuju bentuk yang bisa dieksekusi.

## Boundary

- fokus pada alur kerja compiler Go;
- cocok untuk pembaca yang ingin memahami transformasi kode dari high-level ke low-level;
- bukan tempat utama untuk scheduler atau memory allocator runtime.

---
*Status: [x] Complete*
