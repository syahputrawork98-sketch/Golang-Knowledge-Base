# BK-03: Performance Optimization

Buku ini membahas teknik-teknik tingkat lanjut untuk memeras performa maksimal dari aplikasi Go. Fokusnya adalah pada pemahaman cara kerja compiler (Inlining & BCE), efisiensi penggunaan CPU cache melalui perataan data (Alignment), serta pola desain untuk meminimalkan alokasi memori (Zero-Allocation).

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-inline-bce/README.md) | Inlining & BCE | Compiler flags `-gcflags="-m"`, Leaf functions, Bounds Check Elimination |
| [CH-02](./CH-02-alignment-padding/README.md) | Data Alignment | Struct padding, Word alignment (8-byte), Memory footprint reduction |
| [CH-03](./CH-03-zero-alloc/README.md) | Zero-Allocation | sync.Pool, Object reuse, Reset pattern, GC pressure reduction |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/inlining-logic.svg` | Decision logic for function integration vs stack call |
| `CH-02/assets/struct-padding.svg` | Memory layout comparison (Bad vs Good Alignment) |
| `CH-03/assets/pool-lifecycle.svg` | The Get/Put recycling flow in sync.Pool |

---
*Back to [SR-05 Page](../README.md)*
