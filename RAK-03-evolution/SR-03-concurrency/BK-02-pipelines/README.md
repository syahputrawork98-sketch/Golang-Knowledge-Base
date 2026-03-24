# BK-02: Pipeline & Flow Patterns

Buku ini membahas pola arsitektur aliran data (data flow) konkuren di Go. Dari distribusi beban ke banyak worker hingga penghasil data yang efisien secara memori.

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-fan-out-in/README.md) | Fan-In & Fan-Out | Worker Pool, Multiplexing, Poison Pill |
| [CH-02](./CH-02-orchestration/README.md) | Pipeline Orchestration | Context Propagation, Stage Isolation, Graceful Teardown |
| [CH-03](./CH-03-generators/README.md) | Generator Patterns | Lazy Evaluation, Infinite Streams, Closure-based Generators |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/fan-out-in.svg` | Industrial worker pool pipeline visualization |
| `CH-02/assets/pipeline-orchestration.svg` | Multi-stage context termination flow |
| `CH-03/assets/lazy-generator.svg` | On-demand streaming architecture |

---
*Back to [SR-03 Page](../README.md)*
