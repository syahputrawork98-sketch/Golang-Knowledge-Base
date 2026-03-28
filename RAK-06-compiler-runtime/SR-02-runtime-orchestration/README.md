# SR-02: Runtime Orchestration

Sub-rak ini membahas bagaimana runtime Go mengatur eksekusi program: scheduler, stack, dan garbage collection.

## Struktur

### [BK-01-scheduler](./BK-01-scheduler/)
Model G-M-P dan cara runtime mengatur goroutine.

### [BK-02-stack](./BK-02-stack/)
Pertumbuhan stack dan implikasinya terhadap eksekusi fungsi.

### [BK-03-garbage-collection](./BK-03-garbage-collection/)
Garbage collector dan cara runtime membersihkan memori yang tidak terpakai.

## Boundary

- fokus pada orkestrasi runtime saat program sedang berjalan;
- relevan untuk pembaca yang ingin memahami perilaku Go di balik goroutine dan memori;
- bukan tempat utama untuk frontend compiler atau cgo.

---
*Status: [x] Complete*
