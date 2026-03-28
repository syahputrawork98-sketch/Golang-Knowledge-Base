# CH-03: Dependency Inversion

## 1. Tahap 1: Source Alignment dan Judul

- **Source Link**: [Go Wiki: Code Review Comments - Interfaces](https://go.dev/wiki/CodeReviewComments#interfaces) | [Effective Go: Interfaces and other types](https://go.dev/doc/effective_go#interfaces_and_types)
- **Framing**: Dalam gaya desain Go, dependency inversion biasanya muncul bukan dari framework besar, tetapi dari interface kecil yang didefinisikan dekat dengan tempat pemakaian.

## 2. Tahap 2: Konsep dan Rasionalitas

### Definisi
Dependency inversion berarti logika tingkat tinggi tidak bergantung langsung pada detail implementasi. Keduanya dipisahkan oleh abstraksi kecil yang mewakili perilaku yang benar-benar dibutuhkan.

### Rasionalitas
Pola ini dipilih karena:

1. **Testing jadi lebih mudah**  
   Dependency nyata bisa diganti dengan mock, stub, atau fake tanpa merombak logika inti.
2. **Detail implementasi lebih mudah diganti**  
   Pergantian database, service client, atau adapter lain tidak harus memaksa perubahan besar pada modul inti.
3. **Coupling menurun**  
   Modul tingkat tinggi fokus pada kebijakan, sedangkan detail teknis ditempatkan sebagai implementasi yang bisa ditukar.

### Analogi Model Mental
Bayangkan fitting lampu di rumah. Instalasi listrik tidak bergantung pada satu merek bohlam tertentu. Selama bohlam memenuhi bentuk konektor yang sesuai, kita bisa menggantinya tanpa membongkar seluruh sistem.

### Terminologi Teknis
- **High-level Module**: bagian yang berisi aturan atau logika utama sistem.
- **Low-level Detail**: implementasi konkret seperti database, adapter, atau client eksternal.
- **Injection**: dependency dikirim dari luar, bukan dibuat diam-diam di dalam modul inti.

## 3. Tahap 3: Visualisasi Sistem

```mermaid
graph TD
    H[High-level Service] --> A[Small Interface]
    D1[Concrete Database] ..|implements|> A
    D2[Mock Database] ..|implements|> A
```

## 4. Tahap 4: Mekanisme Pembuktian

Di Go, prinsip ini terasa alami karena interface tidak perlu diumumkan dari sisi provider. Consumer bisa mendefinisikan interface kecil sesuai kebutuhan aktualnya, lalu menerima implementasi konkret dari luar.

Akibat desain ini:
- service inti tidak tahu detail penyimpanan yang dipakai;
- dependency bisa diganti untuk production atau testing;
- arsitektur tetap ringan karena abstraksi dibuat seperlunya, bukan dibangun menjadi layer besar yang rumit.

## 5. Tahap 5: Lab Praktis

Lihat pembuktian kode di folder [examples/](./examples):
- [02_injected_dependency.go](./examples/02_injected_dependency.go) - Contoh service yang bergantung pada interface sehingga implementasi nyata dan mock bisa saling menggantikan.

---
*Status: [x] Complete*
