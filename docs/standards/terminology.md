# Panduan Terminologi (Go Edition)

Menjaga kesederhanaan dan pragmatisme Go melalui bahasa yang lugas.

## 1. Aturan Penulisan Istilah
- **Direct & Simple**: Jangan gunakan kata-kata yang terlalu kompleks jika ada yang lebih sederhana.
- **No Magic**: Hindari kata "sihir" atau "otomatis" tanpa menjelaskan mekanismenya.

## 2. Senior vs Basic Terms

| Basic Term | Senior Terminology | Konteks |
| :--- | :--- | :--- |
| Fungsi paralel | **Concurrency** | Penggunaan Goroutines. |
| Kirim pesan | **Message Passing** | Komunikasi via Channels. |
| Error handling | **Explicit Error Checking** | Pola `if err != nil`. |
| Interface | **Implicit Satisfaction** | Cara Go mengimplementasi interface. |

## 3. Metode Analogi
- **The Gopher Factory**: Gunakan perumpamaan jalur perakitan pabrik untuk menjelaskan konkurensi.
- **Pipa (Pipes)**: Gunakan analogi pipa untuk menjelaskan aliran data di Channels.
