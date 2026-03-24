# CH-02: Composition vs Inheritance (Embedding Power)

> **Source Link**: [Effective Go: Embedding](https://golang.org/doc/effective_go#embedding)

## 1. Konsep & Esensi (Definisi & Rasionalitas)

### Definisi ("Apa itu?")
Go tidak memiliki pewarisan (Inheritance) tradisional (is-a). Sebagai gantinya, Go menggunakan **Komposisi (Composition)** melalui **Struct Embedding** (has-a), di mana sebuah tipe dapat menyisipkan tipe lain ke dalamnya untuk "mewarisi" metode-metodenya tanpa membentuk hierarki kelas yang kaku.

### Rasionalitas ("Why & How?")
1. **Fragile Base Class Problem**: Di inheritance, perubahan di parent bisa merusak ribuan child. Komposisi lebih stabil karena setiap komponen tetap independen.
2. **Concurrency-Friendly**: Hierarki yang dalam sangat sulit dikelola dalam sistem konkuren. Komposisi yang datar jauh lebih mudah dipahami dan didebug.
3. **Flexibility**: Kita bisa mencampur banyak komponen (traits/mixins) ke dalam satu struct dengan sangat mudah.

### Analogi Model Mental
Bayangkan membangun sebuah **Mobil**. Bukannya membuat "Kendaraan -> Mobil" (Inheritance), kita memasang komponen **Mesin**, **Roda**, dan **Rangka** ke dalam satu unit Mobil. Jika kita ingin membuat mobil listrik, kita tinggal mengganti komponen Mesin dengan **Baterai**. Mesin dan Roda tetaplah Mesin dan Roda; mereka tidak perlu tahu mereka ada di dalam Mobil.

---

## 2. Visualisasi Sistem (Mermaid)

```mermaid
classDiagram
    class Logger {
        Log(msg string)
    }
    class DBHandler {
        Logger
        Query()
    }
    DBHandler *-- Logger : Embeds
    Note for DBHandler "DBHandler.Log() calls embedded Logger.Log()"
```

---

## 3. Mekanisme Pembuktian (Algoritma Detil)
Saat kita memanggil metode pada tipe yang membungkus (outer), Go Compiler akan mencari metode tersebut terlebih dahulu pada tipe outer. Jika tidak ada, compiler akan mencari di level embedding (inner). Proses ini disebut **Method Promotion**.

---

## 4. Lab Praktis (Examples)
Silakan tinjau folder [examples/](./examples) untuk eksperimen berikut:
- `01_embedding_basics.go`: Cara menggunakan struct embedding sederhana.
- `02_promotion_shaddowing.go`: Memahami prioritas metode antara tipe outer vs inner.

---
*Unit ini memenuhi standar Platinum Gold (PPM V4).*
