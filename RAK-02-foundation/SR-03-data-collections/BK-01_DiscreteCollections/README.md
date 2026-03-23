# BK-01: Discrete Collections (Array & Slice)

Buku ini membedah tipe data koleksi linear di Go. Kita akan mempelajari perbedaan mendasar antara Array yang statis dan Slice yang dinamis, serta bagaimana Go mengelola keduanya di level memori.

## 📚 Daftar Bab

### 1. [CH-01: Arrays (The Static Base)](./CH-01_Arrays/)
Mengenal fondasi koleksi Go: Ukuran tetap, *Value Semantics*, dan alokasi *stack*.

### 2. [CH-02: Slice Anatomy (Header)](./CH-02_SliceAnatomy/)
Membongkar *Slice Header* (Pointer, Length, Capacity) untuk memahami referensi memori.

### 3. [CH-03: Dynamic Growth (Append)](./CH-03_DynamicGrowth/)
Mekanisme ekspansi memori dinamis melalui fungsi `append()`.

### 4. [CH-04: Memory Safety (Reslicing)](./CH-04_MemorySafety/)
Teknik *slicing* yang aman dan cara menghindari kebocoran memori pada data besar.

---
*Status: [x] Complete (PPM V4)*
