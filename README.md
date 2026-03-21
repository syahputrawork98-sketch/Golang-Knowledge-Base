# Golang Knowledge Base

> **"Mastering Go: From Simple Syntax to Cloud-Native Scale."**

## 🏛️ Arsitektur 6-Rak (Universal Standard)
Repositori ini menggunakan **6-Rack Universal Architecture** dengan prinsip **Digital Mirroring** untuk memisahkan antara fondasi penggunaan dengan dekonstruksi arsitektur mesin.

```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#00ADD8', 'primaryTextColor': '#FFF'}}}%%
graph TD
    Root["Golang Knowledge Base"]
    
    RAK01["RAK-01-anatomy<br/>(The Landscape)"]
    RAK02["RAK-02-foundation<br/>(The Standard Book)"]
    RAK03["RAK-03-evolution<br/>(History & Future)"]
    RAK04["RAK-04-core-mechanics<br/>(The Internal Logic)"]
    RAK05["RAK-05-standard-library<br/>(The Environment)"]
    RAK06["RAK-06-compiler-runtime<br/>(The Machine Room)"]
    
    Root --> RAK01 & RAK02 & RAK03 & RAK04 & RAK05 & RAK06
    
    style Root fill:#00ADD8,stroke:#333,stroke-width:4px,color:#FFF
    style RAK01 fill:#fff,stroke:#333
    style RAK02 fill:#fff,stroke:#333
    style RAK03 fill:#fff,stroke:#333
    style RAK04 fill:#ddd,stroke:#333
    style RAK05 fill:#fff,stroke:#333
    style RAK06 fill:#ddd,stroke:#333
```

---

## 🗄️ Struktur Perpustakaan

### 1. [RAK-01-anatomy](./RAK-01-anatomy/)
Menelusuri esensi naratif Go, filosofi "Less is More", dan target utamanya.

### 2. [RAK-02-foundation](./RAK-02-foundation/)
Sintaks dasar dan eksekusi instruksi Go bersumber dari dokumentasi resmi.

### 3. [RAK-03-evolution](./RAK-03-evolution/)
Mendokumentasikan evolusi spesifikasi, migrasi rilis, dan pergeseran generasi.

### 4. [RAK-04-core-mechanics](./RAK-04-core-mechanics/)
Mekanika Paling Mendalam: error handling, pointer, interface, dan channels.

### 5. [RAK-05-standard-library](./RAK-05-standard-library/)
Eksplorasi Runtime Environment: paket bawaan `net/http`, `fmt`, `sync`, dll.

### 6. [RAK-06-compiler-runtime](./RAK-06-compiler-runtime/)
Deep dive mutlak ke ruang mesin: Go compiler, Goroutine Scheduler (M:P:G), dan Garbage Collection.

---

## 📏 Standar Kualitas (Gold Standard)
Setiap materi mengikuti prinsip **Digital Mirroring** dan standar **PPM V4** yang mewajibkan:
1. **Source-Synced**: Akurasi 1:1 terhadap dokumentasi resmi/spesifikasi.
2. **Experimental Lab**: Kode pembuktian fungsional di folder `examples/`.
3. **Mental Model Visual**: Diagram Mermaid di folder `assets/`.
4. **Narrative Excellence**: Penjelasan mendalam dengan analogi dunia nyata.

*Dokumentasi Lengkap Standar: [docs/standards/architecture.md](./docs/standards/architecture.md)*

---
*Status Pengembangan: [status.md](./status.md)*