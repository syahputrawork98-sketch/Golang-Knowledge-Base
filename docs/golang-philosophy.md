# Filosofi & Esensi Golang

> **"Less is more. Simplicity is the art of hiding complexity."**

## Mengapa Golang Ada?

Dunia software seringkali terjebak dalam perlombaan menambahkan fitur kompleks. Golang hadir bukan untuk menambah kerumitan, melainkan untuk memberikan **Kesederhanaan yang Skalabel**. Di Go, kita tidak bicara tentang "bagaimana menggunakan fitur bahasa yang canggih", tapi tentang "bagaimana menulis kode yang paling mudah dibaca oleh rekan tim".

## Konsep "The Gopher Factory" 🐹

Go tidak mencoba meniru bahasa lain. Ia memiliki pakemnya sendiri yang unik:

### 1. Composition over Inheritance (Bukan Class)
Go secara radikal membuang konsep hierarki *class inheritance* yang seringkali menjadi benang kusut di proyek besar. Go lebih suka kita menyusun komponen kecil menjadi sistem yang lebih besar (**Composition**). "Anda adalah apa yang Anda bisa lakukan (Behavior), bukan dari mana Anda berasal (Type Origin)."

### 2. CSP Model (Goroutines & Channels)
Filosofi paling terkenal di Go adalah:
> *"Do not communicate by sharing memory; instead, share memory by communicating."*
Go menyediakan **Goroutines** (pekerja ringan) dan **Channels** (ban berjalan) untuk berkomunikasi tanpa perlu rasa takut terhadap *deadlock* atau *race condition* tradisional.

### 3. Pragmatik & Berani Berbeda
Go tidak memiliki `try/catch` atau `map/filter/reduce` (sebelum generics hadir secara terbatas). Go lebih memilih perulangan `for` yang jelas dan penanganan *error* yang eksplisit. Filosofinya: **"Error adalah nilai biasa"**—bukan pengecualian yang muncul tiba-tiba.

## Misi Kita

Misi dari perpustakaan ini adalah mentransformasi Anda dari pengembang yang terbiasa dengan pola pikir tradisional menjadi seorang **Go Architect** yang mampu membangun mesin (sistem) yang efisien, berumur panjang, dan sangat tangguh.

---
*Langkah selanjutnya: [RAK-01-foundations](../README.md#rak-01-foundations-the-gateway)*
