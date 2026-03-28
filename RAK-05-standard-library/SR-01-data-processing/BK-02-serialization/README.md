# BK-02: Serialization

Buku ini membahas format serialisasi yang paling penting di standard library Go: JSON untuk interoperabilitas umum dan binary encoding untuk data berukuran tetap yang lebih dekat ke mesin.

## Struktur

### [CH-01_JSON](./CH-01_JSON/)
`encoding/json` untuk marshaling, unmarshaling, struct tags, dan penanganan data dinamis.

### [CH-02_Binary](./CH-02_Binary/)
`encoding/binary` untuk pembacaan dan penulisan data fixed-size berbasis urutan byte.

## Boundary

- fokus pada serialisasi bawaan Go yang paling umum dipakai engineer aplikasi;
- membantu pembaca memahami trade-off antara format teks dan format biner;
- bukan tempat utama untuk protokol jaringan tinggi atau format file pihak ketiga.

---
*Status: [x] Complete*
