# CH-03: Dependency Inversion (Decoupling for Testability)

> **Source Link**: [Solid Principles in Go](https://dave.cheney.net/2016/08/20/solid-go-design) | [Go Blog: Go Interfaces](https://blog.golang.org/interfaces)

## 1. Konsep & Esensi (Definisi & Rasionalitas)

### Definisi ("Apa itu?")
Dependency Inversion Principle (DIP) menyatakan bahwa modul tingkat tinggi tidak boleh bergantung pada modul tingkat rendah. Keduanya harus bergantung pada abstraksi (Interface).

### Rasionalitas ("Why & How?")
1. **Mockability**: Kita bisa dengan mudah mengganti koneksi database asli dengan **Mock/Stub** saat melakukan Unit Testing.
2. **Maintenance**: Perubahan pada detail implementasi (misal: ganti database dari MySQL ke PostgreSQL) tidak memerlukan perubahan pada logika bisnis inti.
3. **Pluggability**: Memungkinkan sistem plugin di mana komponen baru bisa ditambahkan tanpa mengubah kode utama.

### Analogi Model Mental
Bayangkan **Bongkar Pasang Lampu Bohlam**.
Lampu Bohlam tidak disolder langsung ke kabel listrik rumah. Ada sebuah **Fitting (Interface)**. Anda bisa memasang bohlam merek apa saja (Implementasi) selama pas di fitting tersebut. Jika bohlam mati, Anda tidak perlu membongkar instalasi listrik rumah; cukup ganti bohlamnya.

---

## 2. Visualisasi Sistem (Mermaid)

```mermaid
graph TD
    subgraph HighLevelModule
        UserService
    end
    subgraph Abstraction
        UserRepositoryInterface
    end
    subgraph LowLevelModule
        MySQLRepo
        MockRepo
    end
    UserService --> UserRepositoryInterface
    MySQLRepo ..|> UserRepositoryInterface
    MockRepo ..|> UserRepositoryInterface
```

---

## 3. Mekanisme Pembuktian (Algoritma Detil)
Pola ini di Go sangat kuat karena sifat interface-nya yang implisit. Anda bisa mendefinisikan interface justru di tempat ia akan digunakan (Consumer-side interface), yang secara radikal memutus ketergantungan antar paket.

---

## 4. Lab Praktis (Examples)
Silakan tinjau folder [examples/](./examples) untuk eksperimen berikut:
- `01_hard_dependency.go`: Contoh kode yang sulit ditest (Tanpa DIP).
- `02_injected_dependency.go`: Refactor menggunakan interface untuk memudahkan testing.

---
*Unit ini memenuhi standar Platinum Gold (PPM V4).*
