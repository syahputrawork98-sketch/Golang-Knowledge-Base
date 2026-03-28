# BK-02: Tracing

Buku ini membahas cara melihat alur eksekusi program secara lebih detail lewat execution trace, task semantik, dan distributed tracing lintas layanan.

## Struktur

### [CH-01-tracer](./CH-01-tracer/)
Execution trace untuk melihat event runtime, timeline goroutine, dan alur kerja scheduler.

### [CH-02-user-tasks](./CH-02-user-tasks/)
Task dan region buatan pengguna untuk memberi makna semantik pada trace.

### [CH-03-otel-intro](./CH-03-otel-intro/)
Distributed tracing dasar dengan konteks dan trace propagation antar layanan.

## Boundary

- fokus pada pembacaan alur eksekusi dan konteks observability;
- membantu pembaca menghubungkan perilaku runtime dengan unit kerja aplikasi yang nyata;
- bukan tempat utama untuk runtime internals mendalam atau implementasi observability vendor tertentu.

---
*Status: [x] Complete*
