# Panduan Estetika Visual (Go Edition)

Fokus pada kebersihan dan efisiensi visual (Clean & Pragmatic).

## 1. Skema Warna (Branding)
- **Primary Color**: `#00ADD8` (Gopher Blue).
- **Secondary Color**: `#E0EBF5` (Soft Blue).
- **Neutral**: `#222222` (Industrial Black).

## 2. Standar Mermaid
Diagram harus linear dan mudah dibaca:
```mermaid
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#00ADD8', 'edgeLabelBackground':'#fff'}}}%%
graph LR
    A[Goroutine] -- Channel --> B[Worker]
```

## 3. Simbol Visual
- **Gopher Icon**: Gunakan untuk poin-poin tips atau "Gopher Way".
- **Panah Tebal**: Mewakili aliran data utama.
- **Box Tanpa Border**: Mewakili konsep abstrak.
