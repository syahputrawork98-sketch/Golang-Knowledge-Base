# Sejarah & Evolusi Golang

> **"Software Engineering at Google Scale."**

## 1. Kelahiran (2007 - 2009)

Go (sering disebut Golang) lahir dari rasa frustrasi di Google pada tahun 2007. Proyek ini diinisiasi oleh tiga tokoh legendaris: **Robert Griesemer**, **Rob Pike** (pencipta sistem operasi Plan 9 dan encoding UTF-8), serta **Ken Thompson** (pencipta bahasa B, C, dan sistem operasi UNIX).

### Latar Belakang Masalah
Pada saat itu, Google mengelola basis kode yang sangat masif menggunakan C++ dan Java. Namun, mereka menghadapi tiga masalah besar:
- **Kompilasi Lambat**: Menunggu berjam-jam hanya untuk satu kali build.
- **Kompleksitas Berlebih**: Bahasa yang ada terlalu rumit dan sulit dipelihara oleh tim besar.
- **Era Multicore**: Bahasa lama tidak didesain untuk memanfaatkan hardware modern (multicore) secara efisien tanpa kerumitan *threading* yang tinggi.

## 2. Visi Awal: Simplicity & Productivity

Visi para pencipta Go sangat jelas: Go harus menjadi bahasa yang **simpel**, **cepat dikompilasi**, dan **efisien secara runtime**.

Filosofinya adalah: 
- **Bukan bahasa eksperimental**: Go dibuat untuk menyelesaikan masalah nyata, bukan untuk menambah fitur akademik yang rumit.
- **Static Binary**: Go tidak berjalan di mesin virtual (seperti JVM); ia dikompilasi menjadi satu file binary mandiri yang sangat cepat dijalankan.

## 3. Era Modern & Standar Cloud-Native

Sejak rilis versi 1.0 pada tahun 2012, Go menjadi standar industri untuk infrastruktur modern:
- **Lahirnya Docker & Kubernetes**: Dua teknologi revolusioner di dunia modern ini dibangun sepenuhnya menggunakan Go.
- **Ekosistem Cloud-Native**: Go menjadi bahasa pilihan nomor satu untuk layanan *backend* performa tinggi, *microservices*, dan sistem infrastruktur.
- **Go 1 Promise**: Komitmen stabilitas di mana kode yang ditulis hari ini dijamin tidak akan rusak oleh update versi Go di masa mendatang.

## 4. Kenapa Namanya "Go"?

Nama "Go" dipilih karena singkat dan mencerminkan kecepatan (agility). Penggunaan logo **Gopher** (tikus tanah) yang lucu menjadi maskot ikonik yang melambangkan komunitas yang ramah namun pekerja keras dalam membangun sistem yang solid.

---
*Kembali ke [Halaman Utama](../README.md)*
