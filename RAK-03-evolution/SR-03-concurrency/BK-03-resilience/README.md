# BK-03: Resilience & Flow Control

Buku ini membahas pola-pola penting untuk menjaga ketahanan sistem (system resilience) dan mengatur aliran permintaan (flow control) di bawah tekanan.

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-rate-limiting/README.md) | Rate Limiting | Token Bucket, Burst Capacity, `x/time/rate` |
| [CH-02](./CH-02-errgroup/README.md) | Concurrent Error Handling | `errgroup`, Short-Circuit, Context Propagation |
| [CH-03](./CH-03-singleflight/README.md) | Singleflight | Request Coalescing, Cache Stampede Prevention |

---
*Back to [SR-03 Page](../README.md)*
