# BK-03: Telemetry

Buku ini membahas cara mengumpulkan, mengekspor, dan menyusun data operasional aplikasi Go lewat runtime metrics, Prometheus, dan structured logging.

## Struktur

### [CH-01-metrics](./CH-01-metrics/)
`runtime/metrics` sebagai API modern untuk membaca data internal runtime.

### [CH-02-prometheus](./CH-02-prometheus/)
Prometheus sebagai pola umum untuk ekspor dan scraping metrik aplikasi.

### [CH-03-slog](./CH-03-slog/)
Structured logging dengan `log/slog` untuk observability yang lebih mudah diproses mesin.

## Boundary

- fokus pada telemetry yang bisa langsung dipakai engineer aplikasi;
- membantu pembaca membangun jalur data operasional dari runtime sampai dashboard dan log pipeline;
- bukan tempat utama untuk tracing distributed atau internals metric collector yang terlalu vendor-specific.

---
*Status: [x] Complete*
