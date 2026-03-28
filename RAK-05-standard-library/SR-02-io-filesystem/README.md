# SR-02: IO and Filesystem

Sub-rak ini membahas fondasi aliran data dan interaksi sistem file melalui standard library Go.

## Struktur

### [BK-01-io-abstractions](./BK-01-io-abstractions/)
Reader, writer, dan abstraksi I/O yang menjadi dasar banyak package Go.

### [BK-02-file-management](./BK-02-file-management/)
Operasi file, integrasi dengan OS, dan pemakaian `embed`.

## Boundary

- fokus pada I/O dan filesystem sebagai fondasi sistem dan aplikasi;
- membantu pembaca memahami desain package Go yang banyak bertumpu pada interface I/O;
- bukan tempat utama untuk serialisasi atau web API.

---
*Status: [x] Complete*
