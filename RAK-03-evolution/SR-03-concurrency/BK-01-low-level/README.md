# BK-01: Low-Level Concurrency Primitives

Buku ini membahas primitif concurrency tingkat rendah di Go yang dipakai saat channel saja tidak cukup: eksekusi sekali, object reuse, dan signaling berbasis kondisi.

## Struktur

### [CH-01-sync-once](./CH-01-sync-once/)
`sync.Once` dan turunannya untuk inisialisasi satu kali yang aman.

### [CH-02-sync-pool](./CH-02-sync-pool/)
`sync.Pool` untuk object reuse dan pengurangan tekanan alokasi.

### [CH-03-cond-semaphore](./CH-03-cond-semaphore/)
`sync.Cond` dan semaphore untuk signaling dan pembatasan resource yang lebih eksplisit.

## Boundary

- fokus pada primitive concurrency lanjutan dari sisi engineer aplikasi;
- membantu pembaca memilih alat saat channel bukan solusi yang paling tepat;
- bukan tempat utama untuk scheduler runtime mendalam atau filosofi desain concurrency Go.

---
*Status: [x] Complete*
