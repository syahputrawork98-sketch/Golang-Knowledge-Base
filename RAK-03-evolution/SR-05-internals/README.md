# SR-05: System Internals & Optimization

**Mechanical Sympathy: Engineering for the Metal**

Selamat datang di modul puncak **System Internals & Optimization**. Di sini, Anda akan melampaui sekadar menulis kode yang berfungsi dan mulai menulis kode yang bekerja secara harmonis dengan runtime Go dan arsitektur CPU modern.

## Struktur Modul

Modul ini dibagi menjadi tiga buku eksplorasi mendalam:

### [BK-01: Memory Management Internals](./BK-01-memory/README.md)
Memahami siklus hidup objek dan strategi pembersihan memori.
- Escape Analysis (Stack vs Heap decision)
- Garbage Collection Deep Dive (Mark & Sweep internals)
- Memory Allocator Architecture (mspan/mcache hierarchy)

### [BK-02: Runtime Scheduler (G-M-P)](./BK-02-scheduler/README.md)
Membedah mesin penggerak konkurensi Go.
- G-M-P Model (M:N Scheduling logic)
- Syscall & Network Poller (I/O orchestration)
- Preemption Mechanisms (Handling greedy goroutines)

### [BK-03: Performance Optimization](./BK-03-optimization/README.md)
Teknik memeras setiap siklus CPU yang tersedia.
- Inlining & BCE (Compiler-level boosts)
- Data Alignment (Padding & Cache efficiency)
- Zero-Allocation Patterns (sync.Pool & Reuse)

## Kenapa Ini Penting? (Senior Perspective)
Di skala sistem terdistribusi yang memproses jutaan request per detik, penghematan 8 byte per struct atau pengurangan 1ms pada jeda GC bukan lagi optimasi prematur—itu adalah kebutuhan infrastruktur. Modul ini membekali Anda dengan insting "mekanikal" untuk membangun sistem yang benar-benar berperforma tinggi.

---
*Part of [RAK-03-evolution](../README.md)*
