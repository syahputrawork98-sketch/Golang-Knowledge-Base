# BK-01: Module Lifecycle

Buku ini membahas siklus hidup modul Go dari struktur dasar `go.mod` dan `go.sum`, pengembangan multi-module dengan `go.work`, sampai aturan semver untuk versi mayor.

## Struktur

### [CH-01-mod-anatomy](./CH-01-mod-anatomy/)
Struktur, peran, dan hubungan antara `go.mod`, `go.sum`, serta resolusi dependency.

### [CH-02-go-work](./CH-02-go-work/)
Workflow pengembangan multi-module lokal dengan `go.work`.

### [CH-03-semver](./CH-03-semver/)
Aturan semantic import versioning dan perubahan mayor `v2+`.

## Boundary

- fokus pada lifecycle modul dari perspektif workflow engineer;
- membantu pembaca memahami bagaimana toolchain Go menjaga dependency tetap eksplisit dan stabil;
- bukan tempat utama untuk private proxy enterprise atau audit keamanan supply chain yang lebih lanjut.

---
*Status: [x] Complete*
