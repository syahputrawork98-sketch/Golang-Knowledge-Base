# BK-01: Memory Management Internals

Buku ini membahas mekanisme bagaimana Go mengelola memori di bawah kap. Fokusnya adalah memahami perbedaan antara alokasi Stack dan Heap, bagaimana Garbage Collector bekerja secara konkuren, serta arsitektur allocator berbasis span (TCMalloc).

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-escape-analysis/README.md) | Escape Analysis | Heap vs Stack, `-gcflags="-m"`, Share-up vs Share-down |
| [CH-02](./CH-02-gc-deep-dive/README.md) | Garbage Collection | Tri-color marking, STW, Write Barriers, Pacing |
| [CH-03](./CH-03-allocator/README.md) | Memory Allocator | mspan, mcache, mcentral, mheap architecture |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/escape-logic.svg` | Decision tree for stack vs heap allocation |
| `CH-02/assets/tri-color-marking.svg` | The Black, Grey, and White object states in GC |
| `CH-03/assets/allocator-hierarchy.svg` | Hierarchical memory request flow (Local to Global) |

---
*Back to [SR-05 Page](../README.md)*
