# BK-02: Error Management

> **"Errors are values. Don't just check them, handle them gracefully."**

Buku ini membahas revolusi penanganan error di Go, mulai dari pola *Wrapping* hingga mekanisme pertahanan sistem yang tangguh.

## Daftar Bab

1.  **[CH-01: Error Wrapping](./CH-01_Wrapping/README.md)**
    - Menambah konteks pada error dengan `%w` dan inspeksi via `errors.Is`/`As`.
2.  **[CH-02: Custom Errors](./CH-02_CustomErrors/README.md)**
    - Mendesain tipe error yang kaya data untuk arsitektur besar.
3.  **[CH-03: Defer, Panic & Recover](./CH-03_SafetyNet/README.md)**
    - Mekanisme pertahanan terakhir untuk menjaga ketersediaan sistem.

---
*Status: [x] Complete (Gold Standard - PPM V4)*
