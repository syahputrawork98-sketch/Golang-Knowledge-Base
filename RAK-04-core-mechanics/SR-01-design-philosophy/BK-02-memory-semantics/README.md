# BK-02: Memory Semantics

Buku ini membahas bagaimana Go memandang nilai, pointer, dan tata letak memori sebagai dasar pengambilan keputusan desain kode.

## Struktur

### [CH-01_ValueVsPointer](./CH-01_ValueVsPointer/)
Perbedaan nilai dan pointer, serta dampaknya terhadap ownership, mutasi, dan biaya operasional kode.

### [CH-02_MemoryLayoutAndPadding](./CH-02_MemoryLayoutAndPadding/)
Layout memori, alignment, padding, dan efeknya terhadap efisiensi struktur data.

## Boundary

- fokus pada semantik memori dari sudut pandang desain dan implikasi engineering;
- membantu pembaca memahami keputusan kapan memakai value atau pointer;
- bukan tempat utama untuk allocator runtime atau bedah garbage collector rendah.

---
*Status: [x] Complete*
