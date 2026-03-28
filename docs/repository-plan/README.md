# Golang Knowledge Base Blueprint

Dokumen ini memetakan gambaran besar isi repositori. Fokusnya adalah struktur pembelajaran, bukan aturan teknis mikro.

## 1. Fungsi Dokumen Ini

Dokumen ini menjawab:
- repo ini dibagi ke area apa saja;
- kenapa area itu dipisah;
- arah pertumbuhan repositori ke depan.

Aturan teknis repo hidup di:
- `docs/standards/README.md`
- `docs/standards/authoring.md`

## 2. Blueprint 6 Rak

### RAK-01: Anatomy
Membahas asal-usul, filosofi, use case, dan konteks kelahiran Go.

### RAK-02: Foundation
Membahas fondasi bahasa dan penggunaan idiomatik sehari-hari.

### RAK-03: Evolution
Membahas module system, testing, concurrency patterns, observability, dan area evolusi ekosistem.

### RAK-04: Core Mechanics
Membahas rasionalitas desain dan pola sistemik Go.

### RAK-05: Standard Library
Membahas utilitas dan paket bawaan Go sebagai lingkungan kerja utama.

### RAK-06: Compiler and Runtime
Membahas compiler internals, scheduler, allocator, stack, GC, dan interfacing ke dunia luar.

## 3. Kedalaman Struktur

Repositori ini menggunakan struktur:

1. Root
2. RAK
3. SR
4. BK
5. CH
6. SEC

Pengecualian:
- `RAK-01` boleh melewati `SR` karena sifatnya naratif.

## 4. Catatan Strategis

- `RAK-02`, `RAK-04`, dan `RAK-06` harus dijaga boundary-nya agar tidak membahas hal yang sama dari sudut pandang yang sama.
- Pertumbuhan repo harus mengikuti struktur, bukan kebiasaan menambah folder secara ad hoc.
## 5. Rujukan

- [../standards/README.md](../standards/README.md)
- [../standards/authoring.md](../standards/authoring.md)
