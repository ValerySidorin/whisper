package exporters

import (
	"fmt"

	"github.com/ValerySidorin/whisper/internal/config"
)

type Exporter interface {
	SendMessage(msg string, chatID int64) error
}

func Get(cfg *config.Exporter) (Exporter, error) {
	switch cfg.Provider {
	case "telegram":
		return NewTelegramExporter(cfg.Token)
	default:
		return nil, fmt.Errorf("unknown event provider: %v", cfg.Provider)
	}
}
