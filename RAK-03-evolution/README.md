# RAK-03: Evolution (Advanced Gopher Ecosystem)

> **"From Language User to System Architect."**  
> Rak ini tidak hanya mengajarkan 'fitur', tapi 'ekosistem' dan 'pola desain' yang membuat Go menjadi standar industri Cloud-Native.

---

## 🗺️ Expanded 6-Level Roadmap

### [SR-01: Module System & Supply Chain]
**BK-01: The Module Lifecycle**
- CH-01: Anatomy of `go.mod` & `go.sum` (Requirement graphs)
- CH-02: Workspace Workflow (`go.work`) - Managing multi-repo local dev.
- CH-03: Semantic Versioning (v2+ pathing & major version increments).
**BK-02: Enterprise Proxy & Private Hubs**
- CH-01: GOPROXY & GOSUMDB (Security & Reproducibility).
- CH-02: Private Module Authentication (Git Configs & Netrc).
**BK-03: Maintenance & Security**
- CH-01: `go list` & `go mod edit` (The invisible toolings).
- CH-02: Vulnerability Checking (`govulncheck`).
- CH-03: Version Retraction (Fixing 'Mistakes' in Production).

### [SR-02: Testing & Quality Excellence]
**BK-01: Unit Testing Primitives**
- CH-01: Table-Driven Design (The Idiomatic Way).
- CH-02: Test Helpers & Main Setup (`TestMain`).
- CH-03: Gold-Standard File Testing (Golden Files).
**BK-02: The Art of Mocking**
- CH-01: Interface Injection (The 'Classic' Mock).
- CH-02: Higher-Order Function Mocks.
- CH-03: Real Database Testing (Docker/Testcontainers vs Mocking).
**BK-03: Modern Reliability**
- CH-01: Fuzz Testing (Exploding your logic with random data).
- CH-02: Race Detection in CI/CD.
- CH-03: Coverage Analysis & Enforcement.

### [SR-03: Advanced Concurrency Patterns]
**BK-01: Low-Level Sync Primitives**
- CH-01: `sync.Once` & `sync.OnceFunc` (Safe initialization).
- CH-02: `sync.Pool` (Combating Garbage Collector pressure).
- CH-03: semaphores & `sync.Cond`.
**BK-02: Pipeline & Flow Patterns**
- CH-01: Fan-In & Fan-Out (Industrial Scale).
- CH-02: Pipeline Orchestration (Context propagation in chains).
- CH-03: Generator Patterns.
**BK-03: Resilience & Flow Control**
- CH-01: Rate Limiting (Token Buckets vs Leaky Buckets).
- CH-02: `errgroup` (Concurrent error handling).
- CH-03: `singleflight` (Preventing Cache Stampede).

### [SR-04: Runtime Observability & Optimization]
**BK-01: Profiling & Tracing**
- CH-01: `pprof` (CPU, Memory, Block, Mutex profiling).
- CH-02: `go tool trace` (Visualizing Goroutine latency).
- CH-03: Flame Graphs & Heap Analysis.
**BK-02: Build Meta & Embedded Systems**
- CH-01: LDFlags (Injecting commit/version info).
- CH-02: Build Tags (Conditional compilation for OS/Architecture).
- CH-03: The Go `embed` system (Packaging assets as binary).

### [SR-05: Ecosystem & Production Standards]
**BK-01: Project Layout (Standard Go Project)**
- CH-01: `/cmd` vs `/internal` vs `/pkg` (The modern consensus).
- CH-02: Configuration Management (Env vs Files).
**BK-02: API Design & Documentation**
- CH-01: Functional Options Pattern.
- CH-02: Middleware Design Patterns.
- CH-03: GoDoc & Pkg.go.dev (Standard documentation).

---

## ⚡ Technical Standards
- **Depth**: Setiap Bab (CH) akan memiliki minimal 2-3 Section (SEC) mendalam.
- **Visuals**: Diagram aliran data pprof, relasi dependency graph, dan animasi worker pool.
- **Labs**: Minimal 3 code files per chapter di folder `examples/`.
