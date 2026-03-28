# SR-05: Internals

Sub-rak ini membahas area internal yang membantu engineer memahami performa dan perilaku Go lebih dalam tanpa harus langsung masuk ke rak compiler/runtime.

## Struktur

### [BK-01-memory](./BK-01-memory/)
Escape analysis, garbage collection, dan allocator dari sudut pandang performa dan pemahaman perilaku.

### [BK-02-scheduler](./BK-02-scheduler/)
GMP model, syscall poller, dan preemption sebagai bagian dari perilaku runtime yang penting dipahami.

### [BK-03-optimization](./BK-03-optimization/)
Inlining, alignment, padding, dan teknik optimisasi yang terasa langsung efeknya pada kode.

## Boundary

- fokus pada internals yang berguna untuk engineer dan performance-minded developer;
- menjadi jembatan dari ekosistem praktis ke pemahaman mesin yang lebih dalam;
- bukan tempat utama untuk source code compiler Go secara penuh.

---
*Status: [x] Complete*
