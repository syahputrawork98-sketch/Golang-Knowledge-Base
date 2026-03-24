# [BK-03-CH-02] Prometheus Integration

**Industrial Telemetry Export**
*Target: Mengekspos metrik aplikasi Anda ke dashboard Grafana menggunakan standar Prometheus dalam waktu < 4 menit.*

## 1. Definisi & Konsep (The Logic)

**Prometheus** adalah sistem monitoring berbasis "Pull Model" yang sangat populer di ekosistem Cloud Native. Di Go, kita menggunakan paket `prometheus/client_golang` untuk mendefinisikan metrik custom dan mengeksposnya melalui HTTP endpoint (biasanya `/metrics`).

### Terminologi Utama (Senior Terms)
- **Registry**: Wadah pusat tempat seluruh metrik didaftarkan agar bisa dikumpulkan bersama.
- **Pull Model**: Prometheus Server yang mendatangi aplikasi Anda untuk mengambil data, bukan aplikasi yang mengirim data (Push).
- **Labels (Dimensions)**: Pasangan key-value yang ditambahkan ke metrik untuk memungkinkan filtering dan grouping yang fleksibel (misal: `method="GET"`, `status="200"`).

## 2. Rasionalitas (Why & How?)

Mengapa Prometheus menjadi standar de-facto?
- **Multidimensional Data**: Memungkinkan analisis performa per-endpoint, per-user, atau per-region menggunakan label.
- **Alerting**: Integrasi mudah dengan Alertmanager untuk memberi tahu tim jika suatu metrik (misal: error rate) melewati ambang batas.
- **Visualization**: Menjadi sumber data utama bagi Grafana untuk membuat dashboard monitor real-time.

### Mekanisme Kerja Under-the-Hood
1. Aplikasi mendefinisikan variabel global metrik (seperti `CounterVec` atau `HistogramVec`).
2. Saat event terjadi (misal request selesai), aplikasi memanggil `.Inc()` atau `.Observe()`.
3. Endpoint `/metrics` dijalankan oleh `promhttp.Handler()`, yang mengonversi seluruh data registry ke format teks Prometheus yang bisa dibaca.

## 3. Implementasi Utama (The Lab)

Lihat pembuatan exporter metrik di [examples/](./examples/).
1. `01-prom-exporter`: Menyiapkan HTTP server sederhana yang mengekspos metrik jumlah request dan latensi menggunakan library Prometheus.

## 4. Model Mental Visual (The Assets)

![Prometheus Pull Model](./assets/prom-pull-model.svg)

### Prometheus Pull Architecture
```mermaid
graph LR
    App[Go Application] -- Registry --> Handler[/metrics]
    Prom[Prometheus Server] -- Scrape/Pull Every 15s --> Handler
    Prom -- Query --> Grafana[Grafana Dashboard]
    
    subgraph "Application Logic"
    Worker[Worker Process] -- Record --> App
    end
```

---
*Back to [SR-04 Page](../../README.md)*
