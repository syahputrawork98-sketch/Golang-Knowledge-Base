# CH-04: Context Pattern (Lifecycle Management)

> **Source Link**: [Go Packages: context](https://golang.org/pkg/context/) | [Go Blog: Go Concurrency Patterns: Context](https://blog.golang.org/context)

## 1. Konsep & Esensi (Definisi & Rasionalitas)

### Definisi ("Apa itu?")
`context.Context` adalah interface standar Go untuk membawa sinyal pembatalan (**Cancellation**), batas waktu (**Deadlines**), dan nilai-nilai cakupan permintaan (**Request-scoped values**) melintasi batas API dan antar goroutine.

### Rasionalitas ("Why & How?")
1. **Prevent Resource Leaks**: Memungkinkan kita menghentikan goroutine yang tidak lagi diperlukan (misal: user membatalkan request HTTP).
2. **Reliability**: Menetapkan batas waktu pemrosesan agar sistem tidak menggantung (hanging) selamanya karena kegagalan pihak ketiga (DB/API).
3. **Traceability**: Membawa metadata seperti `request_id` untuk keperluan logging dan tracing antar layanan.

### Analogi Model Mental
Bayangkan sebuah **Perintah Militer**.
Seorang Jenderal (`main`) mengirim regu prajurit (`goroutine`) ke medan perang. Regu tersebut dibekali sebuah **Radio (Context)**. 
- Jika Jenderal mematikan radio (**Cancel**), prajurit harus segera pulang.
- Jika radio berbunyi "Misi selesai dalam 5 menit" (**Timeout**), prajurit harus pulang meski misi belum tuntas.
- Tanpa radio, prajurit mungkin akan bertempur selamanya meskipun markas sudah hancur.

---

## 2. Visualisasi Sistem (Mermaid & SVG)

### Hirarki Context (SVG)
![Visualisasi: Context Propagation Tree](./assets/context_tree.svg)

### Alur Pembatalan (Mermaid)
```mermaid
graph TD

    Root[Background/TODO] --> Parent[WithCancel]
    Parent --> Child1[WithTimeout]
    Parent --> Child2[WithValue]
    Child1 --> G1[Goroutine A]
    Child2 --> G2[Goroutine B]
    Note over Parent: Parent cancel triggers Child1 & Child2
```

---

## 3. Mekanisme Pembuktian (Algoritma Detil)
Context membentuk struktur pohon (**Tree Structure**). Jika sebuah context induk dibatalkan, semua context turunannya secara otomatis ikut dibatalkan melalui channel `Done()`. Pengecekan status context biasanya dilakukan secara kooperatif menggunakan `select { case <-ctx.Done(): return }`.

---

## 4. Lab Praktis (Examples)
Silakan tinjau folder [examples/](./examples) untuk eksperimen berikut:
- `01_context_timeout.go`: Menghentikan proses yang terlalu lama.
- `02_context_values.go`: Mengirim metadata secara aman antar fungsi.

---
*Unit ini memenuhi standar Platinum Gold (PPM V4).*
