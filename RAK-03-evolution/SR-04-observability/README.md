# SR-04: Observability

Sub-rak ini membahas cara melihat perilaku program Go saat berjalan lewat profiling, tracing, metrics, dan telemetry.

## Struktur

### [BK-01-profiling](./BK-01-profiling/)
CPU, memory, block, dan mutex profiling untuk membaca biaya eksekusi program.

### [BK-02-tracing](./BK-02-tracing/)
Execution trace dan distributed tracing untuk memahami aliran kerja runtime dan request.

### [BK-03-telemetry](./BK-03-telemetry/)
Metrics, Prometheus, slog, dan telemetry modern dalam aplikasi Go.

## Boundary

- fokus pada observability dan cara membaca perilaku aplikasi secara nyata;
- membantu engineer menjembatani kode dengan data operasional;
- bukan tempat utama untuk testing atau internals compiler.

---
*Status: [x] Complete*
