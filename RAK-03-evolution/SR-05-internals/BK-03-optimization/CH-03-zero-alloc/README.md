# [BK-03-CH-03] Zero-Allocation Patterns

**Winning the GC War with sync.Pool**
*Target: Mengurangi alokasi heap hingga 0% untuk hot-path aplikasi dalam waktu < 4 menit.*

## 1. Definisi & Konsep (The Logic)

**Zero-Allocation** adalah gaya pemrograman di mana kita menghindari pembuatan objek baru di Heap selama fase kritis (hot-path) aplikasi. Hal ini dicapai dengan mendaur ulang objek lama menggunakan **`sync.Pool`** atau pola desain **Reset**.

### Terminologi Utama (Senior Terms)
- **Object Pooling**: Strategi menyimpan objek yang sudah tidak terpakai dalam penampung (pool) untuk diambil kembali nanti, alih-alih membiarkannya dihapus GC.
- **Hot Path**: Jalur kode yang paling sering dieksekusi (misal: loop utama parser atau handler API).
- **Reset Pattern**: Metode pada struct untuk membersihkan isinya tanpa mengalokasikan memori baru (misal: `buf.Reset()`).
- **Victim Cache**: Mekanisme internal `sync.Pool` di mana objek tidak langsung dihapus saat GC, tapi dipindahkan ke area "korban" untuk satu siklus GC lagi.

## 2. Rasionalitas (Why & How?)

Mengapa Senior Developer terobsesi dengan Zero-Allocation?
- **Extreme Throughput**: Alokasi memori adalah operasi yang relatif mahal. Menghilangkan alokasi berarti CPU bisa fokus 100% pada logika bisnis.
- **GC Silence**: Aplikasi dengan alokasi nol tidak akan pernah memicu Garbage Collection, menghilangkan latensi "Stop The World" sepenuhnya.
- **Predictability**: Konsumsi memori menjadi jauh lebih stabil dan dapat diprediksi di bawah beban tinggi.

### Mekanisme Kerja Under-the-Hood (`sync.Pool`)
1. **Get()**: Mencoba mengambil objek dari pool lokal thread (P). Jika kosong, cari di P lain (steal), atau buat baru lewat fungsi `New`.
2. **Put()**: Mengembalikan objek ke pool. Penting: Selalu `Reset()` data sebelum dimasukkan kembali agar tidak ada kebocoran data (dirty state).
3. **GC Integration**: `sync.Pool` otomatis dibersihkan oleh runtime saat GC berjalan untuk mencegah kebocoran memori permanen.

## 3. Implementasi Utama (The Lab)

Lihat perbandingan performa alokasi di [examples/](./examples/).
1. `01-pool-benchmark`: Gunakan `go test -bench . -benchmem` untuk membandingkan kecepatan dan jumlah alokasi antara pendekatan standar vs menggunakan `sync.Pool`.

## 4. Model Mental Visual (The Assets)

![Pool Lifecycle](./assets/pool-lifecycle.svg)

### sync.Pool Lifecycle
```mermaid
graph LR
    App[Go Request] -- Get() --> Pool{sync.Pool}
    Pool -- "Available?" --> Use[Use Object]
    Pool -- "Empty?" --> New[New() Allocation]
    
    Use -- Reset() & Put() --> Pool
    
    subgraph "GC Phase"
    Pool -- "Sweep" --> Removed[Objects Cleared]
    end
```

---
*Back to [SR-05 Page](../../README.md)*
