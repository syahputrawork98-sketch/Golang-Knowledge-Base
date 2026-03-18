# Golang Knowledge Base: Total Deconstruction Plan

> **Status Spec-Sync**: v1.22 (Last Checked: 2026-03-19)

Dokumen ini adalah cetak biru (blueprint) modular yang mendekomposisi **seluruh** dokumentasi resmi Go ([go.dev/doc](https://go.dev/doc/)) ke dalam arsitektur 5-Rack. Desain ini dibuat agar mudah diperbarui seiring perkembangan rilis Go yang aktif (setiap 6-12 bulan).

---

## 🏗 Justifikasi Teknis (The 5-Rack Logic)

Arsitektur ini didesain berdasarkan **filosofi teknis Go** yang memisahkan antara sintaks, model eksekusi, dan keandalan sistem skala besar. Pemisahan ini mencerminkan struktur internal toolchain dan runtime Go.

### 1. RAK-01: Setup & Onboarding (The Zero-Entry Barrier)
- **Justifikasi**: Go didesain untuk "Get Productive Fast". Berbeda dengan ekosistem lain yang membutuhkan konfigurasi lingkungan yang kompleks, Go mengandalkan **Single Toolchain** (`go` command). Rak ini membedah fondasi toolchain, manajemen dependensi modern (`go.mod`), dan tutorial awal yang menjadi gerbang utama efisiensi pengembang.
- **Sub-Rak (SR)**: Dibagi menjadi 2 (Getting Started & Guided Tutorials) untuk memisahkan antara "Instalasi Fisik" dan "Pengenalan Praktis" (Tour/Tutorials).

### 2. RAK-02: Language Fundamentals (The Pragmatic Heart)
- **Justifikasi**: Inti dari Go adalah simplisitas yang pragmatis. Go secara radikal menghindari kerumitan OOP tradisional (seperti *class inheritance*) dan menggantinya dengan **Composition over Inheritance** (Structs) serta **Implicit Interfaces**. Rak ini membedah "Spirit of Go" yang tertuang dalam *Language Spec* dan *Effective Go*.
- **Sub-Rak (SR)**: Dibagi menjadi 3 (Foundations, Data Structures, Abstraction) mengikuti progresi logis: "Atoms" (Syntax), "Molecules" (Data structures), dan "Systems" (Abstraction/Interfaces).

### 3. RAK-03: Concurrency & Runtime (The Engine/CSP Model)
- **Justifikasi**: Ini adalah distingsi utama (USP) dari Go. Rak ini membedah model **Communicating Sequential Processes (CSP)** melalui Goroutines dan Channels. Berbeda dengan model *Event Loop* pada JS/TS, Go menggunakan *m-n scheduler* (G-M-P model) yang memungkinkan efisiensi konkurensi tingkat tinggi dengan biaya memori minimal.
- **Sub-Rak (SR)**: Dibagi menjadi 3 (Primitives, Sync & Control, Deep Dives) untuk memisahkan alat dasar, manajemen keamanan (Sync/Context), dan mekanisme internal (Memory/Reflection).

### 4. RAK-04: Engineering & Reliability (The Production-Ready Brain)
- **Justifikasi**: Go lahir dari kebutuhan sistem skala besar di Google. Rak ini membedah budaya **Built-in Reliability**. Di Go, testing, benchmarking, dan profiling (`pprof`) bukanlah tambahan (*afterthought*), melainkan fitur inti yang terintegrasi secara native.
- **Sub-Rak (SR)**: Dibagi menjadi 3 (Tooling, Observability, Standards) untuk memisahkan eksekusi (Build/Test), visibilitas (Profiling), dan standardisasi industri (Layout/Docs).

### 5. RAK-05: Ecosystem & Evolutions (The Allies)
- **Justifikasi**: Filosofi **"Go 1 Compatibility Promise"** menjamin stabilitas jangka panjang. Rak ini membedah kekayaan **Standard Library** yang sangat komprehensif (Batteries Included), memungkinkan pembangunan sistem *production-grade* tanpa ketergantungan berlebih pada library pihak ketiga.
- **Sub-Rak (SR)**: Dibagi menjadi 2 (Std Lib Mastery & Evolutions) untuk memisahkan materi yang "Statis/Mature" dengan perkembangan bahasa yang "Dinamis" (Release Logs).

---

## 🗄 Peta Arsitektur Detail

| Rak | Sub-Rak | Buku (BK) | Deskripsi Singkat BK |
| :--- | :--- | :--- | :--- |
| **RAK-01** | SR-01 | BK-01: Installing | Setup environment & GOROOT/GOPATH. |
| | | BK-02: Hello World | Siklus hidup `go run` & struktur awal. |
| | | BK-03: Modules Intro | Konsumsi package pertama kali. |
| | SR-02 | BK-01: Create Module | Inisialisasi dan publisitas dependensi. |
| | | BK-02: Workspace | Multi-module local development (`go.work`). |
| | | BK-03: Fuzz Intro | Automated bug discovery di level fungsional. |
| **RAK-02** | SR-01 | BK-01: Basic Syntax | Variables, Zero Values, & Constants. |
| | | BK-02: Flow Control | `for`, `if`, `switch`, & `defer`. |
| | SR-02 | BK-01: Pointers | Address abstraction & safe manipulation. |
| | | BK-02: Structs | Data aggregation & anonymous fields. |
| | | BK-03: Slices | Deep dive into internal length & capacity. |
| | SR-03 | BK-01: Methods | Menambahkan perilaku via receivers. |
| | | BK-02: Interfaces | Kontrak implisit & Duck Typing ala Go. |
| | | BK-03: Errors | Strategi "Errors as Values" & wrapping. |
| | | BK-04: Generics | Parameterized types & type constraints. |
| **RAK-03** | SR-01 | BK-01: Goroutines | Lightweight thread management. |
| | | BK-02: Channels | Komunikasi sinkron & asinkron. |
| | | BK-03: Select | Multiplexing channel communications. |
| | SR-02 | BK-01: Sync Pkg | RWMutext, WaitGroups, & Atomic ops. |
| | | BK-02: Context | Cancellation propagation & timeouts. |
| | SR-03 | BK-01: Reflection | Tipe data `reflect` & runtime inspection. |
| | | BK-02: Memory | Garbage Collection & Escape Analysis. |
| **RAK-04** | SR-01 | BK-01: Tooling | Bedah mendalam command `build`, `vet`, `mod`. |
| | | BK-02: Testing | Table-driven testing pattern. |
| | SR-02 | BK-01: Profiling | CPU & Memory analysis menggunakan `pprof`. |
| | | BK-02: Tracing | Visualisasi runtime execution timeline. |
| | SR-03 | BK-01: Project Layout | Standarisasi folder `/cmd` & `/pkg`. |
| | | BK-02: Documentation | Menulis docstrings standar `godoc`. |
| **RAK-05** | SR-01 | BK-01: Net/HTTP | Web server & client produksi. |
| | | BK-02: Database/SQL | Driver pattern & connection pooling. |
| | | BK-03: Serialization | JSON, XML, dan Protobuf handling. |
| | SR-02 | BK-01: Release Log | Sejarah evolusi fitur (Go 1.x). |
| | | BK-02: Proposal Env | Proses evolusi bahasa via Gopher way. |

---

## 🔄 Protokol Sinkronisasi (6-Month Sync Rule)

Mengingat rilis Golang yang sangat aktif (rata-rata setiap 6 bulan), repository plan ini wajib mengikuti protokol pembaruan berikut untuk menjaga relevansi materi:

1.  **Audit Bi-Annual**: Setiap rilis minor Go baru (e.g., v1.23, v1.24), `repository-plan` harus diaudit ulang terhadap [Official Go Release Notes](https://go.dev/doc/devel/release).
2.  **Modular Update**: Jika ada fitur baru (seperti *Generics* atau *Fuzzing* di masa lalu), tentukan penempatan Rak & Buku yang paling logis tanpa merombak total struktur yang ada.
3.  **Versioning Log**: Setiap hasil audit harus dicatat dalam **Log Sinkronisasi** di bawah ini, mencantumkan versi target Go dan perubahan arsitektural yang dilakukan.

---

## 📜 Log Sinkronisasi (Spec-Log)

Log ini berfungsi sebagai bukti historis bahwa *Knowledge Base* ini tetap mutakhir dengan spesifikasi bahasa terbaru.

| Versi Go | Tanggal Audit | Perubahan Arsitektur/Konten | Status |
| :--- | :--- | :--- | :--- |
| **v1.22** | 2026-03-19 | Inisialisasi awal 5-Rack & Justifikasi Teknis. | ✅ Synced |
| *v1.23* | *TBD* | *Antisipasi rilis berikutnya...* | ⚪ Planned |

---

## 📈 Prinsip Pengerjaan "Tanpa Setengah-setengah"

1.  **Iterasi Bottom-Up**: Mulai dari Bab (CH) terkecil, selesaikan 3 Tahap PPM V4.
2.  **Referensial**: Setiap Bab WAJIB mencantumkan sumber URL dokumentasi resmi di baris pertama.
3.  **Visual Sentris**: Jangan biarkan ada Buku tanpa diagram "The Factory" di level Rak.
4.  **Continuous Sync**: Selalu periksa Log Sinkronisasi sebelum memulai pengerjaan Bab baru untuk memastikan basis teori masih relevan dengan versi Go target.
