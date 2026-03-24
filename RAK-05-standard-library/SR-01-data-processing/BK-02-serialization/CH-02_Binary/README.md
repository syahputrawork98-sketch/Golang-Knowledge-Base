# CH-02: Binary Encoding (Fixed Length)

> **Source Link**: [Go Packages: encoding/binary](https://golang.org/pkg/encoding/binary/)

## 1. Konsep & Esensi (Definisi & Rasionalitas)

### Definisi ("Apa itu?")
Pakat `encoding/binary` menangani konversi antara angka (uint32, int64, dsb) dan urutan byte tetap, seringkali digunakan untuk format file biner atau protokol tingkat rendah.

### Rasionalitas ("Why & How?")
1. **Deterministic Size**: Berbeda dengan JSON yang ukurannya fleksibel (teks), binary encoding menjamin ukuran data yang tetap dan sangat kecil.
2. **Endianness Control**: Memungkinkan pemilihan `BigEndian` atau `LittleEndian` untuk kompatibilitas antar arsitektur CPU yang berbeda.
3. **Speed**: Konversi langsung ke byte jauh lebih cepat daripada konversi ke teks dan kembali lagi.

### Analogi Model Mental
Bayangkan **Mesin Morse**.
JSON adalah bahasa Inggris (bisa panjang-pendek). **Binary** adalah kode Morse pendek-titik-garis yang sangat efisien dikirim lewat kabel. Pakat `binary` adalah **Operator Morse** yang memastikan kode tersebut ditulis dengan urutan yang benar (ujung besar/kecil) agar bisa dibaca di seberang sana.

---

## 2. Visualisasi Sistem (Mermaid)

```mermaid
graph LR
    I[Int64: 1024] -->|binary.Write| B[[]byte: [0 0 4 0]]
    B -->|binary.Read| I2[Int64: 1024]
    
    subgraph CPU_Order
        BE[Big Endian: MSB First]
        LE[Little Endian: LSB First]
    end
```

---

## 3. Mekanisme Pembuktian (Algoritma Detil)
Pakar menyarankan penggunaan `binary.Read` dan `binary.Write` hanya untuk data dengan ukuran tetap (*fixed-size values*). Go menggunakan refleksi untuk mengidentifikasi tipe data, namun untuk kecepatan ekstrem, gunakan fungsi spesifik seperti `PutUint32` atau `Uint32` dari `LittleEndian/BigEndian` secara langsung pada slice byte.

---

## 4. Lab Praktis (Examples)
Silakan tinjau folder [examples/](./examples) untuk eksperimen berikut:
- `01_read_binary_file.go`: Membaca header file biner (seperti gambar/audio).
- `02_endianness_proof.go`: Membuktikan perbedaan posisi byte antara arsitektur berbeda.

---
*Unit ini memenuhi standar Platinum Gold (PPM V4).*
