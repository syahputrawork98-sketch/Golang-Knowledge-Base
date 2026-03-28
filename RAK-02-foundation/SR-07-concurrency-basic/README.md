# SR-07: Concurrency Basic

Sub-rak ini membahas fondasi concurrency Go: bagaimana goroutine bekerja, bagaimana channel dipakai, dan bagaimana kode tetap aman saat berjalan bersamaan.

## Struktur

### [BK-01_Orchestration](./BK-01_Orchestration/)
Goroutine, scheduling dasar, dan orkestrasi pekerjaan paralel.

### [BK-02_Communication](./BK-02_Communication/)
Channel dan pola komunikasi antar goroutine.

### [BK-03_Safety](./BK-03_Safety/)
Mutex, context, dan dasar menjaga kode concurrent tetap aman.

## Boundary

- fokus pada dasar concurrency yang wajib dipahami sebelum masuk ke pola lanjutan;
- menyiapkan pembaca untuk rak evolusi dan core mechanics;
- bukan tempat utama untuk concurrency pattern industri atau bedah scheduler runtime.

---
*Status: [x] Complete*
