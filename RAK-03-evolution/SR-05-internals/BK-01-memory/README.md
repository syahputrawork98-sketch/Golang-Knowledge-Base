# BK-01: Memory Internals

Buku ini membahas bagaimana Go mengelola memori dari sudut pandang engineer: escape analysis, perilaku garbage collector, dan allocator runtime yang memengaruhi performa nyata.

## Struktur

### [CH-01-escape-analysis](./CH-01-escape-analysis/)
Escape analysis untuk memahami kapan data tetap di stack dan kapan dipaksa naik ke heap.

### [CH-02-gc-deep-dive](./CH-02-gc-deep-dive/)
Garbage collector Go dari perspektif kerja konkuren, pacing, dan biaya latensi.

### [CH-03-allocator](./CH-03-allocator/)
Allocator runtime untuk membaca jalur permintaan memori dari cache lokal sampai heap global.

## Boundary

- fokus pada internals memori yang berguna untuk membaca perilaku performa aplikasi;
- membantu pembaca menjembatani kode sehari-hari dengan keputusan runtime di bawahnya;
- bukan tempat utama untuk bedah source runtime Go secara penuh.

---
*Status: [x] Complete*
