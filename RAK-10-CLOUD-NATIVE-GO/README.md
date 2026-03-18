# Rak 10: Cloud-Native Go

Fokus pada peran Go sebagai bahasa utama dalam infrastruktur cloud modern (The Backbone of Cloud).

## Struktur Sub-Rak
- **[SR-01-CONTAINER-DYNAMICS](./SR-01-CONTAINER-DYNAMICS/README.md)**: Integrasi Go dengan Docker & K8s Internals.
- **[SR-02-SERVERLESS-GO](./SR-02-SERVERLESS-GO/README.md)**: Go dalam AWS Lambda & Edge Computing.

## Go and the Cloud
Go diciptakan di Google untuk menyelesaikan masalah "Scale". Karakteristiknya membuatnya sangat cocok untuk Cloud:
1. **Static Binary**: Aplikasi Go dikompilasi menjadi satu file, memudahkan deployment di Docker (Image size kecil).
2. **Fast Startup**: Sangat krusial untuk lingkungan Serverless/Auto-scaling.
3. **Memory Efficiency**: Menangani ribuan koneksi dengan RAM minimal.

---
*Kembali ke [Documentation Hub](../docs/README.md)*
