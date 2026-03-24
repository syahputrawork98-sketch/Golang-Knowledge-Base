# BK-01: Low-Level Sync Primitives

Buku ini membahas komponen sinkronisasi tingkat rendah (low-level) di Go yang berada di bawah abstraksi channel. Pemahaman mendalam terhadap primitif-primitif ini adalah tanda seorang Senior Go Engineer.

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-sync-once/README.md) | sync.Once & sync.OnceFunc | Lazy Initialization, Atomic Fast-Path, Go 1.21 Primitives |
| [CH-02](./CH-02-sync-pool/README.md) | sync.Pool | GC Pressure Reduction, Per-P Storage, Work Stealing |
| [CH-03](./CH-03-cond-semaphore/README.md) | Semaphores & sync.Cond | Fan-out Signaling, Wait Queue, Broadcast vs Signal |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/sync-once-internal.svg` | Atomic fast-path & mutex slow-path diagram |
| `CH-02/assets/sync-pool-arch.svg` | Per-Processor storage & work stealing architecture |
| `CH-03/assets/sync-cond-orchestration.svg` | Broadcast signaling & wait queue mechanism |

---
*Back to [SR-03 Page](../README.md)*
