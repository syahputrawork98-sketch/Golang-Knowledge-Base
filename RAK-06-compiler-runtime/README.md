# RAK-06: Compiler and Runtime

Rak ini adalah area terdalam di repo: bagaimana compiler Go bekerja, bagaimana scheduler mengatur goroutine, bagaimana memori dikelola, dan bagaimana Go berinteraksi dengan dunia luar di level sistem.

## Struktur

### [SR-01-compiler-architecture](./SR-01-compiler-architecture/)
Parser, AST, SSA, dan code generation dalam pipeline compiler Go.

### [SR-02-runtime-orchestration](./SR-02-runtime-orchestration/)
Scheduler, stack, dan orkestrasi runtime saat program berjalan.

### [SR-03-memory-management](./SR-03-memory-management/)
Allocator, spans, arenas, dan cara Go mengelola memori secara efisien.

### [SR-04-system-interfacing](./SR-04-system-interfacing/)
Cgo, assembly, dan batas interaksi antara Go dengan sistem atau bahasa lain.

## Boundary

- fokus pada implementasi rendah dan perilaku mesin di balik program Go;
- cocok untuk pembaca yang ingin memahami Go lebih dalam dari sekadar pemakaian API;
- bukan tempat utama untuk pengantar sintaks atau navigasi standard library sehari-hari.

---
*Status: [x] Complete*
