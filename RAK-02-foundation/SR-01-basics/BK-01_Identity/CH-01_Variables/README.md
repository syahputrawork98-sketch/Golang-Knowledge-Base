# CH-01: Variables (Storage & Identity)

> **"Explicit types are documentation; inferred types are convenience."**

---

## 1. Tahap 1: Source Alignments & Judul
- **Source Link**: [Go Spec: Variables](https://go.dev/ref/spec#Variables)

---

## 2. Tahap 2: Konsep & Esensi

### Definisi ("Apa itu?")
Variabel adalah lokasi penyimpanan memori yang memiliki nama, tipe data, dan nilai. Go adalah bahasa yang **Statically Typed**, di mana setiap variabel memiliki tipe data yang kaku dan ditentukan pada waktu kompilasi (*compile-time*).

### Rasionalitas ("Why & How?")
- **Memory Safety**: Go menjamin bahwa setiap variabel memiliki nilai awal yang valid (*Zero Value*), sehingga tidak ada celah untuk mengeksekusi data "sampah" dari sisa memori sistem.
- **Redundancy Reduction**: Dengan *Short Variable Declaration* (`:=`), Go mengurangi kebisingan kode tanpa mengorbankan keamanan tipe data (*type safety*).

### Analogi Model Mental
**Label Kontainer**. Bayangkan variabel sebagai label yang Anda tempelkan pada sebuah kotak penyimpanan di gudang (RAM). Jika Anda melabeli kotak tersebut sebagai "Buku", Anda tidak bisa memasukkan "Cairan" ke dalamnya. Label ini membantu Anda menemukan barang dengan cepat tanpa harus memeriksa setiap inci lantai gudang.

### Terminologi Teknis
- **Lexical Scope**: Batas wilayah di mana sebuah nama variabel dikenal oleh compiler.
- **Zero Value**: Keadaan awal yang deterministik bagi variabel yang dideklarasikan tanpa nilai awal.

---

## 3. Tahap 3: Visualisasi Sistem

### High-Level Model (Mermaid)
```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#00ADD8', 'primaryTextColor': '#FFF'}}}%%
graph LR
    subgraph RAM
        VarA[Address: 0x101 | Value: 10]
        VarB[Address: 0x102 | Value: 'Hello']
    end
    NameA[var x int] --> VarA
    NameB[y := 'Hello'] --> VarB
```

---

## 4. Tahap 4: Mekanisme Pembuktian (Runtime Allocation)

Bagaimana Go menangani variabel di level mesin?
- **Static Resolution**: Saat kompilasi, nama `x` atau `y` akan dihapus dan digantikan dengan offset alamat memori relatif (e.g., `SP+8`).
- **Initialization**: Go Runtime memastikan bahwa saat fungsi dipanggil, area memori yang akan ditempati variabel dibersihkan (*zeroed out*) dalam satu operasi instruksi mesin. Ini membuat Go sangat aman dari bug memori acak yang sering ditemukan pada bahasa C/C++.

---

## 5. Tahap 5: Multi-file Lab Praktis (Examples)

Melihat berbagai cara deklarasi dan perilaku Zero Values secara nyata.

- **Lab 1**: [01_declarations.go](./examples/01_declarations.go) - Eksplorasi sintaks `var` vs `:=`.
- **Lab 2**: [02_zero_values.go](./examples/02_zero_values.go) - Melihat nilai default setiap tipe data.

---
*Status: [x] Complete (Gold Standard - PPM V4)*
