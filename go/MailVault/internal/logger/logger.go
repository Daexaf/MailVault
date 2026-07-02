package logger

import (
	"log/slog"
	"os"
)

// New Membuat logger baru menggunakan slog
func New() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}
