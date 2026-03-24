# BK-02: Runtime Scheduler (G-M-P)

Buku ini membahas orkestrasi internal bagaimana Go menjadwalkan ribuan goroutine di atas thread sistem operasi secara efisien. Fokusnya adalah pada model G-M-P, penanganan I/O yang memblokir, serta mekanisme preemption untuk menjaga keadilan distribusi CPU.

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-gmp-model/README.md) | G-M-P Model | Processor (P), OS Thread (M), Work Stealing, M:N Scheduling |
| [CH-02](./CH-02-syscall-poller/README.md) | Syscall & Netpoller | Sinkron vs Asinkron I/O, Netpoller (epoll), Syscall Handoff |
| [CH-03](./CH-03-preemption/README.md) | Preemption | Cooperative vs Asynchronous, SIGURG signaling, sysmon |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/gmp-architecture.svg` | Hierarchical relationship between G, M, and P |
| `CH-02/assets/io-management.svg` | Decision tree for Network Wait vs File Syscall |
| `CH-03/assets/preemption-signal.svg` | The sysmon to M signaling flow for greedy goroutines |

---
*Back to [SR-05 Page](../README.md)*
