# SR-03: Advanced Concurrency Patterns

**From Go Channel Basics to High-Performance System Architect**

Modul ini membekali Anda dengan senjata-senjata terberat di arsenal Go untuk membangun sistem konkuren yang efisien, aman, dan tahan banting.

## Struktur Modul

| Book | Topik | Target |
|------|-------|--------|
| [BK-01](./BK-01-low-level/README.md) | Low-Level Sync Primitives | Menguasai `sync.Once`, `sync.Pool`, dan `sync.Cond` |
| [BK-02](./BK-02-pipelines/README.md) | Pipeline & Flow Patterns | Merancang aliran data industrial dengan Fan-Out/In dan Generator |
| [BK-03](./BK-03-resilience/README.md) | Resilience & Flow Control | Melindungi sistem dengan Rate Limiting, errgroup, dan Singleflight |

## Learning Path

```mermaid
graph LR
    BK01[BK-01: Low-Level Primitives] --> BK02[BK-02: Pipelines]
    BK02 --> BK03[BK-03: Resilience]
    BK03 --> NEXT[SR-04: Runtime Observability]
```

---
*Back to [RAK-03 Page](../README.md)*
