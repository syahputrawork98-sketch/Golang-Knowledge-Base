# CH-02: Inspection & Switching (Assertion & Switch)

> **"Interface values are dynamic containers. To use the concrete data inside, you must safely open the container."**

---

## 1. Tahap 1: Source Alignments & Judul
- **Source Link**: [Go Spec: Type Assertions](https://go.dev/ref/spec#Type_assertions)
- **Status**: [x] Platinum Gold Standard

---

## 2. Tahap 2: Konsep & Esensi

### Definisi ("Apa itu?")
**Type Inspection** adalah proses untuk mengidentifikasi atau mengekstrak tipe data konkrit yang tersimpan di dalam variabel interface saat program berjalan. Ini dilakukan melalui dua mekanisme: **Type Assertion** (untuk satu tipe) dan **Type Switch** (untuk banyak tipe).

### Rasionalitas ("Why & How?")
- **Granular Access**: Interface hanya memberikan akses ke method yang didefinisikan dalam kontraknya. Jika Anda butuh mengakses field atau method spesifik milik struct asli, Anda harus melakukan "Downcasting" (kembali ke tipe aslinya).
- **Runtime Branching**: Memungkinkan kode untuk berperilaku berbeda berdasarkan tipe data yang masuk tanpa harus menggunakan Generic (misal: memproses input yang bisa berupa `string`, `int`, atau `[]byte`).
- **Safety**: Go menyediakan pola `comma-ok` untuk memastikan proses ekstraksi tidak menyebabkan program crash (*panic*) jika tipenya tidak cocok.

### Analogi Model Mental
**Pemeriksaan Keamanan Bandara**. Interface adalah penumpang yang memakai jaket tebal. Petugas keamanan (Type Switch) tidak tahu siapa di balik jaket itu sampai mereka memeriksa ID-nya. Jika ID-nya "Pilot" (Assertion), dia boleh masuk ke kokpit (Akses method spesifik). Jika ID-nya "Penumpang", dia harus duduk di kursi biasa.

### Terminologi Teknis
- **Comma-OK Idiom**: Pola `v, ok := i.(T)` untuk ekstraksi aman.
- **Type Switch Shadowing**: Pola `switch v := i.(type)` di mana `v` memiliki tipe yang berbeda di setiap `case`.

---

## 3. Tahap 3: Visualisasi Sistem

### Type Switching Flow
```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#00ADD8', 'primaryTextColor': '#FFF'}}}%%
graph TD
    Input[Interface Value] --> Switch{switch i.(type)}
    Switch -->|case int| P1[v is Integer]
    Switch -->|case string| P2[v is String]
    Switch -->|case *User| P3[v is Pointer to User]
    Switch -->|default| P4[Unknown Type]
```

---

## 4. Tahap 4: Mekanisme Pembuktian (Reliability & Performance)

Bagaimana cara melakukan inspeksi yang efisien?
- **Avoid Over-Assertion**: Melakukan assertion berulang kali dalam loop besar bisa berdampak pada performa karena melibatkan pemeriksaan `_type` descriptor. Gunakan Type Switch jika ada lebih dari 2 tipe yang mungkin.
- **The any Trap**: Jangan gunakan `any` di mana-mana hanya agar bisa melakukan Type Switch. Gunakan interface yang bermakna jika memungkinkan untuk menjaga *Compile-time Type Safety*.
- **Panic Protection**: Selalu gunakan pola `ok` saat melakukan assertion tunggal. Baris kode seperti `s := i.(string)` tanpa pengecekan adalah bom waktu yang menunggu untuk meledak saat program menerima input yang salah.

---

## 5. Tahap 5: Multi-file Lab Praktis (Examples)

Mempraktikkan teknik inspeksi tingkat tinggi.

- **Lab 1**: [01_safe_assertion.go](./examples/01_safe_assertion.go) - Teknik ekstraksi tipe tanpa risiko panic.
- **Lab 2**: [02_dynamic_switch.go](./examples/02_dynamic_switch.go) - Menangani multi-tipe dengan pola shadowing.
- **Lab 3**: [03_interface_conversion.go](./examples/03_interface_conversion.go) - Mengubah satu interface ke interface lainnya (Interface to Interface assertion).

---
*Status: [x] Complete (Gold Standard - PPM V4)*
