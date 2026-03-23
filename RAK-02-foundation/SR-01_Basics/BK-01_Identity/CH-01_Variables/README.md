# CH-01: Variables (Storage & Identity)

> **"Explicit types are documentation; inferred types are convenience."**

---

## 1. Tahap 1: Source Alignments & Judul
- **Source Link**: [Go Spec: Variables](https://go.dev/ref/spec#Variables)
- **Analogi**: **Label Kontainer**. Variabel adalah label yang kita tempelkan pada sebuah kotak (memori) untuk menyimpan nilai tertentu. Di Go, label ini sangat ketat: sekali kotak tersebut diberi label "Angka", ia tidak bisa diisi "Teks".

---

## 2. Tahap 2: Concept & Essence

### Apa itu Variabel di Go?
Variabel adalah lokasi penyimpanan memori yang memiliki nama dan tipe data. Go adalah bahasa yang **Statically Typed**, artinya tipe data ditentukan saat kompilasi.

### Why & How?
- **Zero Values**: Di Go, tidak ada variabel yang benar-benar "kosong" atau unitialized. Jika Anda tidak memberi nilai, Go akan memberikan nilai default (e.g., `0` untuk int, `""` untuk string). Ini mencegah bug *NULL pointer* atau memori sampah.
- **Short Variable Declaration (`:=`)**: Digunakan untuk deklarasi cepat di dalam fungsi dengan *type inference*.

### Terminologi Teknis
- **Scope**: Area di mana variabel dapat diakses (Package level vs Local level).
- **Type Inference**: Kemampuan compiler untuk menebak tipe data berdasarkan nilai yang diberikan.

---

## 3. Tahap 3: Visualisasi Sistem (Memory)

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

## 4. Tahap 4: Mekanisme Pembuktian (Zero Values)

Mengapa Go memiliki *Zero Values*? 
- Di bahasa lama seperti C, variabel yang tidak diisi akan mengambil nilai apapun yang tertinggal di memori tersebut (dangerous).
- Go secara otomatis membersihkan memori saat alokasi.
- **Detail Teknis**: Saat `var i int` dipanggil, runtime memastikan bit-bit di lokasi memori tersebut di-set ke nol. Ini memberikan keamanan deterministik sejak awal.

---

## 5. Tahap 5: Multi-file Lab Praktis (Examples)

Melihat berbagai cara deklarasi dan perilaku Zero Values.

- **Lab 1**: [01_declarations.go](./examples/01_declarations.go) - Eksplorasi sintaks `var` vs `:=`.
- **Lab 2**: [02_zero_values.go](./examples/02_zero_values.go) - Melihat nilai default setiap tipe data.

---
*Status: [x] Complete (Gold Standard)*
