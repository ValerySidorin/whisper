package messenger

import (
	"fmt"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger/telegram"
)

func GetExporter(cfg *config.Exporter) (port.Exporter, error) {
	switch cfg.Provider {
	case "telegram":
		return telegram.NewExporter(cfg)
	default:
		return nil, fmt.Errorf("unknown messenger: %v", cfg.Provider)
	}
}
