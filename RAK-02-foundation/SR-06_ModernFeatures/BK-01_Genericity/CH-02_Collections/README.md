# CH-02: Type-safe Collections

> **"Custom data structures no longer need to choose between type safety and code reuse. With Generics, you have both."**

---

## 1. Tahap 1: Source Alignments & Judul
- **Source Link**: [Go Blog: Why Generics?](https://go.dev/blog/why-generics)
- **Status**: [x] Platinum Gold Standard

---

## 2. Tahap 2: Konsep & Esensi

### Definisi ("Apa itu?")
**Type-safe Collections** adalah struktur data kustom (seperti Stack, Queue, Linked List, atau Tree) yang menggunakan **Generic Type Parameters** agar dapat menampung tipe data apa pun tanpa perlu melakukan *Type Assertion* saat mengambil data kembali.

### Rasionalitas ("Why & How?")
- **Eliminate `any` (interface{})**: Sebelum Go 1.18, koleksi umum harus menggunakan `interface{}`. Ini memaksa pengguna untuk melakukan casting `v.(int)` yang bisa menyebabkan panic jika salah.
- **Improved Performance**: Koleksi generic bekerja langsung pada tipe data konkrit. Ini berarti tidak ada overhead *Boxing/Unboxing* atau alokasi heap tambahan yang biasanya terjadi pada interface.
- **Readable APIs**: Dengan generics, definisi koleksi menjadi jauh lebih jelas: `stack := Stack[int]{}`. Siapa pun yang membaca kode tahu persis apa isi stack tersebut.

### Analogi Model Mental
**Rak Buku Berlabel**.
- Tanpa Generics: Anda punya rak umum yang bisa diisi apa saja. Saat butuh Obeng, Anda harus merogoh rak dan berharap yang Anda ambil bukan Palu (Runtime risk).
- Dengan Generics: Anda memesan "Rak Khusus Alat Tulis" (Instantiation). Rak ini hanya menerima Pulpen dan Pensil. Saat Anda mengambil barang, Anda DIJAMIN mendapatkan alat tulis (Compile-time guarantee).

### Terminologi Teknis
- **Generic Struct**: Struct yang didefinisikan dengan type parameter, misal `type Box[T any] struct { content T }`.
- **Receiver with Type Parameter**: Method pada generic struct harus menyertakan type parameter dalam receiver-nya, misal `func (b *Box[T]) Get() T`.

---

## 3. Tahap 3: Visualisasi Sistem

### Generic Stack Architecture
```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#00ADD8', 'primaryTextColor': '#FFF'}}}%%
graph TD
    S[Stack[T]] --> D[Items: []T]
    S --> P[Push val T]
    S --> O[Pop returns T]
    
    T1[T = int] --> S1[Stack[int]]
    T2[T = string] --> S2[Stack[string]]
```

---

## 4. Tahap 4: Mekanisme Pembuktian (Zero Value & Generics)

Masalah kritis saat membuat koleksi generic:
- **Returning Zero Value**: Bagaimana jika kita memanggil `Pop()` pada stack kosong? Kita harus mengembalikan "Zero Value" dari `T`.
- **Solution**: Karena kita tidak tahu tipe `T`, kita bisa mendeklarasikan variabel `var zero T` dan mengembalikannya. Variabel tersebut otomatis akan berisi `0` untuk int, `""` untuk string, atau `nil` untuk pointer/interface.
- **Constraints in Structs**: Kita juga bisa membatasi isi koleksi kita. Misal: `type NumberList[T Number] struct { ... }` hanya akan bisa dibuat untuk tipe angka.

---

## 5. Tahap 5: Multi-file Lab Praktis (Examples)

Membangun struktur data generic yang fungsional.

- **Lab 1**: [01_generic_stack.go](./examples/01_generic_stack.go) - Implementasi Stack (LIFO) yang type-safe.
- **Lab 2**: [02_generic_list.go](./examples/02_generic_list.go) - Implementasi sederhana Generic Linked List.

---
*Status: [x] Complete (Gold Standard - PPM V4)*
