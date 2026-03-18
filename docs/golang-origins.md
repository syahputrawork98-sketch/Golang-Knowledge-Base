# Asal-usul & Filosofi Golang: "Efisiensi Skala Google"

> **"Sebuah jawaban atas rasa frustrasi menunggu proses kompilasi C++ yang tak kunjung usai."**

Meskipun `golang-history.md` memberikan garis waktu, dokumen ini menggali lebih dalam tentang narasi di balik penciptaan Go.

---

## 📅 Kilas Balik: Frustrasi di Kantor Google (2007)
Kisah ini dimulai di sebuah ruang kerja di Google. Tiga tokoh legendaris—**Robert Griesemer**, **Rob Pike**, dan **Ken Thompson**—sedang menunggu sebuah proyek C++ raksasa dikompilasi. Proses itu memakan waktu begitu lama sehingga mereka memiliki waktu untuk mendiskusikan betapa rumit dan tidak efisiennya bahasa pemrograman yang mereka gunakan saat itu.

Mereka menyadari bahwa bahasa-bahasa besar pada masa itu dirancang sebelum era komputasi awan (*cloud*), prosesor multi-core, dan basis kode berskala jutaan baris.

---

## 💡 Masalah yang Ingin Diselesaikan
Go diciptakan untuk menyelesaikan "penyakit" infrastruktur di Google:

1.  **Slow Builds**: Menunggu berjam-jam untuk satu kali perubahan kecil sangat menghambat produktivitas.
2.  **Uncontrolled Complexity**: C++ dan Java menjadi terlalu kompleks dengan terlalu banyak fitur "ajaib" yang sulit dipahami oleh pengembang baru.
3.  **Concurrency Crisis**: Bahasa lama tidak memiliki cara yang simpel dan aman untuk memanfaatkan ribuan *core* prosesor secara bersamaan.

---

## 🏆 Kehebatan Golang: Tiga Pilar Utama

### 1. "Less is More"
Go secara radikal membuang fitur-fitur yang dianggap menambah kerumitan tanpa manfaat sebanding (seperti *inheritance*, *assertions*, dan penanganan pengecualian `try/catch`). Go memilih kesederhanaan agar kode sangat mudah dibaca oleh siapa pun.

### 2. Kompilasi Secepat Kilat
Go dirancang agar bisa dikompilasi dalam hitungan detik, bukan jam. Ini memberikan siklus pengembangan yang sangat cepat (*Edit-Compile-Run*).

### 3. Concurrency as First-Class Citizen
Dengan **Goroutines** dan **Channels**, Go membuat pemrograman paralel menjadi sangat murah (cepat dan hemat memori) serta aman. "Jangan berkomunikasi dengan berbagi memori; berbagi memorilah dengan berkomunikasi."

---

## 🎭 Analogi: "Kereta Api Cepat yang Minimalis"
Jika C++ adalah mesin uap raksasa yang sangat kuat namun membutuhkan ribuan tuas dan awak untuk dioperasikan, maka **Go** adalah **Kereta Api Cepat (Shinkansen)** yang didesain aerodinamis.

Ia hanya memiliki sedikit tuas kendali yang esensial, namun ia bisa melaju jauh lebih cepat dengan efisiensi energi yang luar biasa tinggi untuk mengangkut data dalam jumlah masif.

---

## 🚀 Para Pendiri (The Fathers)
- **Ken Thompson**: Pencipta sistem operasi UNIX dan bahasa B (leluhur C).
- **Rob Pike**: Tokoh utama sistem operasi Plan 9 dan UTF-8.
- **Robert Griesemer**: Arsitek engine V8 JavaScript dan Java HotSpot VM.

---
*Kembali ke [Halaman Utama](../README.md)*
