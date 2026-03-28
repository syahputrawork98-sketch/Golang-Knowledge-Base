# BK-01: Unit Testing

Buku ini membahas fondasi unit testing idiomatik di Go: table-driven tests, helper dan lifecycle setup, serta golden files untuk output yang kompleks.

## Struktur

### [CH-01-table-driven](./CH-01-table-driven/)
Table-driven tests sebagai pola dasar pengujian yang ringkas dan mudah diperluas.

### [CH-02-helpers-setup](./CH-02-helpers-setup/)
Penggunaan helper, `TestMain`, dan `t.Cleanup` untuk menjaga lifecycle test tetap rapi.

### [CH-03-golden-files](./CH-03-golden-files/)
Golden files untuk memvalidasi output besar tanpa mengotori test dengan string raksasa.

## Boundary

- fokus pada pola dasar yang membangun confidence dalam unit test Go;
- membantu pembaca menulis test yang idiomatik, mudah dirawat, dan jelas dibaca;
- bukan tempat utama untuk mocking lanjut atau tooling reliabilitas seperti fuzzing dan race detector.

---
*Status: [x] Complete*
