# Golang Knowledge Base: Total Deconstruction Plan

> **Status Spec-Sync**: v1.24 (Full Alignment)
> **Last Updated**: 2026-03-19

Arsitektur **Source-Driven 9-Rack** ini mencerminkan taksonomi asli [go.dev/doc](https://go.dev/doc/).

---

## 🏗 Justifikasi Teknis (The Mirroring Principle)

Setiap Rak dipetakan langsung ke kategori navigasi utama di dokumentasi resmi Go.

### 1. RAK-01: Getting Started
Instalasi, Hello World, dan pengaturan workspace.

### 2. RAK-02: Learning Go
Panduan naratif utama (User Manual, Effective Go, Go Tour).

### 3. RAK-03: Language Specification
Aturan formal bahasa Go.

### 4. RAK-04: The Standard Library
Dokumentasi paket bawaan (API Reference).

### 5. RAK-05: Modules Reference
Manajemen dependensi dan sistem `.mod`.

### 6. RAK-06: Diagnostics & Tooling
Kompilasi, debugging, profiling, dan tracing.

### 7. RAK-07: Engineering Practices
Testing, bencharking, dan layout proyek standar.

### 8. RAK-08: The Go Memory Model
Teori sinkronisasi dan aturan visibilitas memori (Unique Pillar).

### 9. RAK-09: Evolution & Proposals
Sejarah rilis, changelog, dan proposal fitur masa depan.

---

## 🗄 Peta Arsitektur Detail

| Rak | Sub-Rak (SR) | Buku (BK) | Deskripsi BK |
| :--- | :--- | :--- | :--- |
| **RAK-01** | SR-01: Setup | BK-01: Installation | Install & OS specific setup. |
| | | BK-02: Workspace | `GOPATH` vs Modules mode. |
| **RAK-02** | SR-01: Narrative | BK-01: The Go Tour | Intro interaktif. |
| | | BK-02: Effective Go | Menulis kode idiomatik. |
| | SR-02: User Manual | BK-03: Ref Manual | Panduan penggunaan praktis. |
| **RAK-03** | SR-01: Spec | BK-01: Core Spec | Aturan sintaksis & runtime. |
| **RAK-04** | SR-01: Essentials | BK-01: Built-ins | Types, constants, functions. |
| | SR-02: Packages | BK-02: Core Pkgs | `fmt`, `os`, `io`, `net`. |
| **RAK-05** | SR-01: Modules | BK-01: Go Mod Ref | `go.mod` & `go.sum` logic. |
| **RAK-06** | SR-01: CLI | BK-01: Go Command | Build, install, run utilities. |
| | SR-02: Analysis | BK-02: Profiling | `pprof` & Performance metrics. |
| | | BK-03: Tracing | Execution tracer. |
| **RAK-07** | SR-01: QA | BK-01: Testing | `testing` package & logic. |
| | | BK-02: Benchmarking | Performance testing. |
| | SR-02: Design | BK-03: Project Layout | Standard project structures. |
| **RAK-08** | SR-01: Theory | BK-01: Sync Rules | Happens-before relationship. |
| **RAK-09** | SR-01: Pulse | BK-01: Release Notes | Changelog v1.0 to v1.24. |
| | | BK-02: Proposals | Design docs for new features. |

---

## 📜 Log Sinkronisasi (Spec-Log)

| Versi Go | Tanggal Audit | Perubahan Arsitektur | Status |
| :--- | :--- | :--- | :--- |
| **v1.24** | 2026-03-19 | Migrasi ke 9-Rack (Full Source Mirroring). | ✅ Synced |
