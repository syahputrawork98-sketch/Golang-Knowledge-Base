# CH-01: Implicit Satisfaction

## 1. Tahap 1: Source Alignment dan Judul

- **Source Link**: [Go Specification: Interface types](https://go.dev/ref/spec#Interface_types)
- **Framing**: Di Go, tipe tidak perlu mendaftarkan diri secara resmi untuk memenuhi interface. Kalau bentuk perilakunya cocok, interface itu otomatis terpenuhi.

## 2. Tahap 2: Konsep dan Rasionalitas

### Definisi
Implicit satisfaction berarti sebuah tipe dianggap memenuhi interface jika ia memiliki seluruh method yang diminta interface tersebut. Tidak ada kata kunci `implements` seperti di banyak bahasa lain.

### Rasionalitas
Desain ini penting karena:

1. **Coupling antar paket jadi lebih rendah**  
   Tipe konkret tidak perlu tahu interface mana saja yang akan memakainya.
2. **Fokus bergeser ke perilaku**  
   Yang penting bukan garis keturunan tipe, tetapi apa yang bisa dilakukan tipe itu.
3. **Komposisi jadi lebih alami**  
   Satu tipe bisa memenuhi banyak interface kecil tanpa deklarasi tambahan.

### Analogi Model Mental
Bayangkan stopkontak dan perangkat listrik. Perangkat tidak perlu mendaftar ke produsen stopkontak. Selama bentuk stekernya cocok, perangkat itu bisa langsung dipakai. Di Go, interface bekerja dengan logika kecocokan perilaku seperti itu.

### Terminologi Teknis
- **Implicit Satisfaction**: pemenuhan interface tanpa deklarasi eksplisit.
- **Structural Typing**: kecocokan berdasarkan bentuk perilaku.
- **Behavior Contract**: kontrak perilaku yang dinyatakan oleh interface.

## 3. Tahap 3: Visualisasi Sistem

```mermaid
classDiagram
    class Reader {
        <<interface>>
        Read(p []byte) (n int, err error)
    }
    class File {
        Read(p []byte) (n int, err error)
    }
    class NetworkBuffer {
        Read(p []byte) (n int, err error)
    }
    File ..|> Reader : satisfies implicitly
    NetworkBuffer ..|> Reader : satisfies implicitly
```

## 4. Tahap 4: Mekanisme Pembuktian

Secara desain, implicit satisfaction membuat interface bisa didefinisikan di sisi pemakai, bukan harus di sisi pembuat tipe konkret. Itu membuat package lebih mudah dipisah dan diuji.

Di level implementasi, runtime interface Go memang punya representasi internal seperti pasangan tipe dan data, tetapi pelajaran utamanya di `RAK-04` adalah: model ini sengaja dipilih untuk menjaga sistem tetap modular dan tidak terseret ke hierarki yang kaku.

## 5. Tahap 5: Lab Praktis

Lihat pembuktian kode di folder [examples/](./examples):
- [01_structural_typing.go](./examples/01_structural_typing.go) - Membuktikan pemenuhan interface tanpa deklarasi `implements`.

---
*Status: [x] Complete*
