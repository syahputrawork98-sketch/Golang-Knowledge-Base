# BK-01: String and Text Processing

Buku ini membahas paket standard library yang paling sering dipakai untuk bekerja dengan teks: manipulasi string, konversi nilai ke dan dari string, serta pengolahan `[]byte` saat data perlu lebih fleksibel.

## Struktur

### [CH-01_Strings](./CH-01_Strings/)
Operasi pencarian, pemotongan, penggabungan, dan builder untuk teks UTF-8 yang umum dipakai sehari-hari.

### [CH-02_Strconv](./CH-02_Strconv/)
Konversi aman antara representasi string dan tipe dasar seperti `int`, `float`, dan `bool`.

### [CH-03_Bytes](./CH-03_Bytes/)
Pemakaian `bytes` dan `bytes.Buffer` saat data perlu dimodifikasi sebagai `[]byte`.

## Boundary

- fokus pada pengolahan teks dan representasi data ringan dengan package bawaan;
- membantu pembaca memilih kapan memakai `string`, `strconv`, atau `[]byte`;
- bukan tempat utama untuk format serialisasi kompleks seperti JSON atau binary protocol.

---
*Status: [x] Complete*
