package messenger

import (
	"fmt"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger/telegram"
)

func Register(cfg *config.Configuration, r port.MessageRenderer) (port.Messenger, error) {
	switch cfg.Messenger.Provider {
	case "telegram":
		return telegram.RegisterTelegram(&cfg.Messenger, r)
	default:
		return nil, fmt.Errorf("unknown messenger: %v", cfg.Messenger.Provider)
	}
}
