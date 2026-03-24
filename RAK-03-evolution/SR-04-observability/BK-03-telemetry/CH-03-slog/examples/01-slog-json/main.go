package main

import (
	"log/slog"
	"os"
)

func main() {
	// 1. Inisialisasi JSON Handler dengan opsi custom
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		// Opsi: Sembunyikan field waktu untuk demo agar output bersih
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	
	// 2. Set Logger default
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// 3. Info Log sederhana
	slog.Info("Application started", "version", "1.1.0")

	// 4. Log dengan Group (Sangat berguna untuk context request)
	logger.Info("Incoming request",
		slog.Group("request",
			slog.String("method", "GET"),
			slog.String("path", "/api/user"),
			slog.Int("status", 200),
		),
		slog.Duration("latency", 150),
	)

	// 5. Error Log (Atribut khusus)
	slog.Error("Database connection failed", "error", "connection timeout")
}
