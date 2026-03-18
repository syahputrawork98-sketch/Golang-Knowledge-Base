# Golang Knowledge Base: Documentation Hub

Pusat dokumentasi ini menyimpan seluruh cetak biru, standar kualitas, dan panduan meta untuk membangun repository ini secara konsisten.

## Daftar Dokumen Utama

1.  **[Architecture Standards](./standards/architecture.md)**: Definisi analogi "The Gopher Factory" dan aturan pewajiban file per tingkatan.
2.  **[Repository Plan](./repository-plan/README.md)**: Dekomposisi total dari dokumentasi resmi Go (`go.dev/doc`) ke dalam struktur 9-Rack disertai justifikasi teknis modular.
3.  **[Golang Origins](./golang-origins.md)**: Kisah di balik layar mengapa Google menciptakan bahasa Go.
4.  **[Golang History](./golang-history.md)**: Sejarah kelahiran Go dari rasa frustrasi di Google hingga era cloud-native.
5.  **[Philosophy & Essence](./golang-philosophy.md)**: Membedah filosofi "Less is More" dan model CSP.
6.  **[Why Golang?](./why-golang.md)**: Rasionalitas teknis kapan harus memilih Go untuk sistem skala besar.

## Struktur Direktori `docs/`

```text
/docs
├── standards/
│   └── architecture.md    # Aturan main/Gold Standard
├── repository-plan/
│   └── README.md          # Cetak biru mapping konten (Modular)
└── README.md              # File ini (Hub navigasi)
```

## Cara Berkontribusi/Update
Setiap kali ada materi baru yang selesai (PPM V4), pastikan untuk memperbarui status di file `status.md` di root directory.
