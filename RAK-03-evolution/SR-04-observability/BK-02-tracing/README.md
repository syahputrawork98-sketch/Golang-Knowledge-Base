# BK-02: Tracing & Execution Flow

Buku ini membahas cara memvisualisikan orkestrasi goroutine dan aliran eksekusi program secara real-time. Fokusnya adalah memahami interaksi antar-goroutine, latensi penjadwal (scheduler), dan bagaimana menghubungkan trace lokal ke sistem terdistribusi.

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-tracer/README.md) | Execution Tracer | STW events, Processor timeline, Handoff visualization |
| [CH-02](./CH-02-user-tasks/README.md) | User-defined Tasks | Semantic spans, context propagation, logical grouping |
| [CH-03](./CH-03-otel-intro/README.md) | Distributed Tracing | OpenTelemetry, TraceID, Context propagation (Network) |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/execution-trace.svg` | Goroutine timeline and scheduler orchestration |
| `CH-02/assets/semantic-trace.svg` | Hierarchy of Tasks and Regions for business logic |
| `CH-03/assets/dist-tracing.svg` | End-to-end distributed tracing pipeline across services |

---
*Back to [SR-04 Page](../README.md)*
