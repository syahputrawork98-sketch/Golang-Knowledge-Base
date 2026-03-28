# BK-03: Resilience and Flow Control

Buku ini membahas pola ketahanan sistem concurrent di Go: pembatasan laju, short-circuit berbasis error, dan penggabungan permintaan duplikat.

## Struktur

### [CH-01-rate-limiting](./CH-01-rate-limiting/)
Rate limiting untuk menjaga sistem tetap stabil saat trafik naik atau permintaan datang terlalu cepat.

### [CH-02-errgroup](./CH-02-errgroup/)
`errgroup` untuk koordinasi banyak goroutine yang harus gagal atau berhenti bersama.

### [CH-03-singleflight](./CH-03-singleflight/)
`singleflight` untuk mencegah cache stampede dan menekan kerja duplikat.

## Boundary

- fokus pada flow control dan resilience pattern di level engineer aplikasi;
- membantu pembaca menjaga sistem concurrent tetap stabil di bawah tekanan;
- bukan tempat utama untuk retry policy besar, observability produksi, atau runtime internals.

---
*Status: [x] Complete*
