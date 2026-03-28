# BK-01: Profiling

Buku ini membahas cara membaca biaya eksekusi program Go lewat `pprof`, profiling goroutine, dan analisis contention agar bottleneck nyata bisa terlihat dari data, bukan tebakan.

## Struktur

### [CH-01-pprof](./CH-01-pprof/)
CPU dan heap profiling untuk menemukan hot path dan pola alokasi yang mahal.

### [CH-02-goroutines](./CH-02-goroutines/)
Profil goroutine dan thread untuk mendeteksi leak serta pola blocking yang tidak sehat.

### [CH-03-contention](./CH-03-contention/)
Block dan mutex contention untuk membaca biaya sinkronisasi di bawah load.

## Boundary

- fokus pada observability performa dari sisi engineer aplikasi;
- membantu pembaca menemukan bottleneck nyata melalui profile dan laporan runtime;
- bukan tempat utama untuk bedah scheduler atau internals allocator yang lebih cocok di rak lain.

---
*Status: [x] Complete*
