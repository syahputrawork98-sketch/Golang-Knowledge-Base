# BK-01: IO Abstractions

Buku ini membahas abstraksi I/O inti di standard library Go: kontrak `Reader` dan `Writer`, serta buffering yang membuat aliran data lebih efisien.

## Struktur

### [CH-01_IO](./CH-01_IO/)
Konsep dasar `io.Reader`, `io.Writer`, dan composability antar sumber serta tujuan data.

### [CH-02_Bufio](./CH-02_Bufio/)
`bufio` untuk buffering, scanning, dan pengurangan syscall saat membaca atau menulis data kecil secara berulang.

## Boundary

- fokus pada fondasi I/O yang dipakai banyak package standard library;
- membantu pembaca memahami kenapa banyak API Go dibangun di atas kontrak `Reader` dan `Writer`;
- bukan tempat utama untuk operasi file OS tingkat tinggi atau serialisasi data.

---
*Status: [x] Complete*
