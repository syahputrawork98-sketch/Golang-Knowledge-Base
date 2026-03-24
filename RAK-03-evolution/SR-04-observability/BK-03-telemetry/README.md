# BK-03: Metrics & Structured Logging

Buku ini membahas cara mengumpulkan data operasional (telemetry) dari aplikasi Go menggunakan standar modern. Fokusnya adalah pada efisiensi pengumpulan metrik runtime, integrasi dengan ekosistem monitoring industri (Prometheus), dan penerapan logging terstruktur yang dapat diproses oleh mesin.

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-metrics/README.md) | runtime/metrics | Modern telemetry API, Low overhead, Histogram semantics |
| [CH-02](./CH-02-prometheus/README.md) | Prometheus Integration | Pull model, Registry, Metrics exposition (/metrics) |
| [CH-03](./CH-03-slog/README.md) | Structured Logging | slog (Go 1.21+), JSON Handlers, Attributes & Groups |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/metric-semantics.svg` | Counter, Gauge, and Histogram data models |
| `CH-02/assets/prom-pull-model.svg` | Scraper architecture and time-series propagation |
| `CH-03/assets/slog-pipeline.svg` | From log record to serialized JSON output pipeline |

---
*Back to [SR-04 Page](../README.md)*
