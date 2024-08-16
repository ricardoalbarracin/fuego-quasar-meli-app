package infraestructure

import (
	"fuego-quasar-app/internal/core/domain/port"
	"log/slog"
	"os"
)

func NewLog() port.LogService {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	return logger
}
