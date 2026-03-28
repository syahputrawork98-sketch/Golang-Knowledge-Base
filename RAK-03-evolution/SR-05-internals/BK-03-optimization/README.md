# BK-03: Performance Optimization

Buku ini membahas optimisasi yang terasa langsung di kode Go: keputusan compiler, tata letak data, dan pengurangan alokasi yang menekan kerja garbage collector.

## Struktur

### [CH-01-inline-bce](./CH-01-inline-bce/)
Inlining dan bounds check elimination untuk memahami optimisasi yang dilakukan compiler sebelum kode berjalan.

### [CH-02-alignment-padding](./CH-02-alignment-padding/)
Alignment dan padding untuk membaca biaya layout struct terhadap memori dan cache.

### [CH-03-zero-alloc](./CH-03-zero-alloc/)
Pola zero-allocation dan object reuse saat performa perlu dijaga tanpa membebani heap.

## Boundary

- fokus pada optimisasi yang relevan untuk engineer aplikasi dan performance-minded developer;
- membantu pembaca memahami mengapa layout kode dan data bisa mengubah performa secara nyata;
- bukan tempat utama untuk membahas source code compiler Go secara penuh.

---
*Status: [x] Complete*
