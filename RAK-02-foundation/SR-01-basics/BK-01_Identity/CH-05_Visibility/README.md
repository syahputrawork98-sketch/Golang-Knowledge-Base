# CH-05: Visibility (The Capitalization Rule)

> **"In Go, the first letter of a name determines its visibility to other packages."**

---

## 1. Tahap 1: Source Alignments & Judul
- **Source Link**: [Go Spec: Exported Identifiers](https://go.dev/ref/spec#Exported_identifiers)

---

## 2. Tahap 2: Konsep & Esensi

### Definisi ("Apa itu?")
**Visibility** (Visibilitas) di Go mengatur apakah sebuah pengenal (fungsi, tipe, variabel) dapat diakses dari paket lain. Go menggunakan mekanisme sederhana: kapitalisasi huruf pertama.

### Rasionalitas ("Why & How?")
- **Automatic Encapsulation**: Menghilangkan kebutuhan untuk mengetik kata kunci seperti `public` atau `private`.
- **API Surface Management**: Membantu pengembang untuk memisahkan logika internal paket (implementasi) dengan antarmuka publik yang aman untuk digunakan oleh paket lain.

### Analogi Model Mental
**Pintu VIP vs Pintu Belakang**. Nama dengan huruf BESAR (Kapital) adalah pintu VIP: tamu dari luar rumah (paket lain) dipersilahkan masuk. Nama dengan huruf kecil adalah pintu belakang: hanya penghuni di dalam rumah (paket yang sama) yang boleh menggunakannya.

### Terminologi Teknis
- **Exported**: Dapat diakses secara publik (Kapital).
- **Unexported**: Terbatas hanya untuk internal paket (Huruf Kecil).

---

## 3. Tahap 3: Visualisasi Sistem

### High-Level Model (Mermaid)
```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#00ADD8', 'primaryTextColor': '#FFF'}}}%%
graph LR
    subgraph Package_External
        Call[main.go]
    end
    subgraph Package_Utility
        Exported[CalcData - Visible]
        Internal[calcInternal - Hidden]
    end
    Call --> Exported
    Call --X Internal
```

---

## 4. Tahap 4: Mekanisme Pembuktian (Symbol Exporting)

Bagaimana Go memaksakan aturan visibilitas?
- **Symbol Table Filter**: Selama proses kompilasi, Go compiler membangun tabel simbol untuk setiap paket. Hanya simbol yang diawali huruf besar yang akan disertakan dalam file objek biner agar dapat "dilihat" (linked) oleh paket lain.
- **Compile-time Enforcement**: Jika paket luar mencoba memanggil `calcInternal`, compiler akan segera menolak dengan error "cannot refer to unexported name", memastikan keamanan sistem bahkan sebelum program dijalankan.

---

## 5. Tahap 5: Multi-file Lab Praktis (Examples)

Mengenal cara kerja enkapsulasi paket.

- **Lab 1**: [visibility_demo.go](./examples/visibility_demo.go) - Eksperimen teori visibilitas antar paket.

---
*Status: [x] Complete (Gold Standard - PPM V4)*
