# SR-04: Runtime Observability

**Mastering the Science of System Introspection**

Selamat datang di modul **Runtime Observability**. Di sini, Anda akan belajar cara melihat ke dalam "jantung" aplikasi Go yang sedang berjalan untuk mengoptimalkan performa, mendeteksi kebocoran resource, dan memantau kesehatan sistem di skala produksi.

## Struktur Modul

Modul ini dibagi menjadi tiga buku utama yang mencakup spektrum observabilitas penuh:

### [BK-01: Profiling & Deep Inspection](./BK-01-profiling/README.md)
Teknik menggunakan `pprof` untuk bedah performa mendalam.
- CPU & Heap Analysis (Bottleneck identification)
- Goroutine & Thread Profiling (Leak detection)
- Block & Mutex Contention (Sync synchronization)

### [BK-02: Tracing & Execution Flow](./BK-02-tracing/README.md)
Visualisasi orkestrasi internal Go Scheduler.
- Execution Tracer (Timeline visualization)
- User-defined Tasks & Regions (Semantic spans)
- Distributed Tracing Concepts (OpenTelemetry)

### [BK-03: Metrics & Structured Logging](./BK-03-telemetry/README.md)
Telemetri modern untuk sistem yang skalabel.
- `runtime/metrics` (Efficient runtime introspection)
- Prometheus Integration (Industrial monitoring)
- Structured Logging with `slog` (Machine-readable logs)

## Kenapa Ini Penting? (Senior Perspective)
Seorang Senior Engineer tidak menebak mengapa aplikasi lambat; mereka **mengukur**. Kemampuan menggunakan profil dan trace secara efektif adalah pembeda utama antara pengembang yang "berharap kodenya cepat" dan yang "tahu kodenya optimal".

---
*Part of [RAK-03-evolution](../README.md)*
