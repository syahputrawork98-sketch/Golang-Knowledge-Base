# Arsitektur & Hierarki Struktur (Go Edition)

Proyek **Golang Knowledge Base** disusun dengan analogi **Pabrik Gopher (The Gopher Factory)** untuk mentransformasi dokumentasi teknis (seperti Go Spec, Effective Go, & standard library) menjadi unit pelajaran yang manusiawi dan efisien.

## Analogi Struktur

Berikut adalah pemetaannya ke dalam direktori bertingkat:

| Tingkatan | Analogi | Contoh Direktori | Keterangan |
| :--- | :--- | :--- | :--- |
| **Level 1** | **Pabrik (Factory)** | `/` (root) | Seluruh sistem proyek (Go Engine). |
| **Level 2** | **Rak (Shelf)** | `RAK-01-foundations/` | Pengelompokan besar domain ilmu. |
| **Level 3** | **Sub-Rak (Sub-shelf)** | `SR-01_Intro/` | Grup materi berbasis area spesifik. |
| **Level 4** | **Buku (Book)** | `BK-01_GoTutorial/` | Koleksi bab yang membentuk satu topik besar. |
| **Level 5** | **Bab (Chapter)** | `CH-01_Overview/` | Unit terkecil wajib (Folder Bab). |

---

## Aturan Pewajiban `README.md`

Guna memudahkan navigasi, setiap tingkatan direktori **WAJIB** memiliki file `README.md`:

- **Root (`/README.md`)**: Visi keseluruhan (The Efficient Gopher 🐹).
- **Rak (`RAK-XX/README.md`)**: Tujuan dan cakupan Rak tersebut.
- **Buku (`BK-XX/README.md`)**: Sinopsis dan orientasi filosofis materi.
- **Bab (`CH-XX/README.md`)**: Materi inti (PPM Stage 1).

---

## Kriteria "Gold Standard" (100% Complete)

Sebuah unit dianggap **Completed** jika memenuhi 4 pilar kualitas:

> [!IMPORTANT]
> **Pilar 1: Spec-Sync Accurate**  
> Konten diverifikasi terhadap dokumentasi resmi/Go Spec (Target di kolom Spec-Sync).
>
> **Pilar 2: Functional Examples**  
> Minimal 1 contoh kode (`.go`) di folder `examples/`.  
> *Catatan*: Jika materi bersifat teoretis/konseptual tanpa kode eksekusi, tetap buat file `.go` dengan komentar: `"// Ini bukan kode, ini hanya materi [NamaBab]"`.
>
> **Pilar 3: Mental Model Visual**  
> Minimal 1 diagram (Mermaid/SVG) di folder `assets/`.  
> *Catatan*: Jika materi benar-benar tidak memerlukan visualisasi, buat placeholder/catatan: `"Tidak memerlukan gambar/SVG narasi untuk [NamaBab]"`.
>
> **Pilar 4: Narrative Excellence**  
> Penjelasan menggunakan standar PPM V4: Manusiawi, ada analogi, dan menggunakan **Bahasa Arsitek (Senior Terminology)**. Dokumentasi harus berfungsi sebagai jembatan yang memungkinkan pengembang awam berkomunikasi secara profesional dengan pengembang senior setelah mempelajari materi.
