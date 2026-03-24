# CH-02: Accept Interfaces, Return Structs (The Go Mantra)

> **Source Link**: [Go Wiki: Interface Design](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)

## 1. Konsep & Esensi (Definisi & Rasionalitas)

### Definisi ("Apa itu?")
Ini adalah prinsip desain API paling fundamental di Go:
- **Accept Interfaces**: Parameter fungsi harus bertipe interface agar fungsi tersebut fleksibel menerima berbagai tipe data (Abstraksi Input).
- **Return Structs**: Fungsi harus mengembalikan tipe konkret (struct) agar pemanggil (caller) memiliki kontrol penuh dan performa maksimal (Konkret Output).

### Rasionalitas ("Why & How?")
1. **Flexibility**: Memungkinkan ketergantungan dibalik (Dependency Inversion). Kita bisa mengganti implementasi input tanpa mengubah fungsi.
2. **Predictability**: Caller tahu persis apa yang mereka dapatkan. Mereka bisa memanggil metode apa saja pada struct tersebut tanpa harus menebak interface mana yang dikembalikan.
3. **Decoupling**: Menghindari pembuatan interface "prematur" sebelum benar-benar dibutuhkan.

### Analogi Model Mental
Bayangkan sebuah **Stopkontak (Function)**.
- **Input**: Berbentuk interface lubang steker universal. Ia siap menerima beban listrik apa pun (Seterika, Kulkas, Lampu) asalkan kakinya cocok.
- **Output**: Adalah **Energi Listrik Konkret (Watt)**. Anda tidak ingin mendapatkan "Sesuatu yang mirip energi"; Anda butuh Watt yang nyata untuk menyalakan perangkat Anda.

---

## 2. Visualisasi Sistem (Mermaid)

```mermaid
graph LR
    I[Input: io.Reader] --> F[Function: ProcessData]
    F --> O[Output: *Report Struct]
    Note over I: Accept Interfaces (Flexible)
    Note over O: Return Structs (Concrete)
```

---

## 3. Mekanisme Pembuktian (Algoritma Detil)
Menerima interface memungkinkan compiler melakukan pengecekan kepuasan tipe di sisi pemanggil. Mengembalikan struct memungkinkan compiler melakukan optimasi alokasi (seperti inlining) yang lebih baik karena tipe data sudah diketahui secara pasti saat compile time.

---

## 4. Lab Praktis (Examples)
Silakan tinjau folder [examples/](./examples) untuk eksperimen berikut:
- `01_flexible_input.go`: Fungsi yang menerima `fmt.Stringer` untuk fleksibilitas log.
- `02_concrete_return.go`: Alasan kenapa mengembalikan struct mempermudah dokumentasi dan penggunaan API.

---
*Unit ini memenuhi standar Platinum Gold (PPM V4).*
