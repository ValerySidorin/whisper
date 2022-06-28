package telegram

import (
	"github.com/ValerySidorin/whisper/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramExporter struct {
	Bot     *tgbotapi.BotAPI
	ChatIDs []int
}

func NewExporter(cfg *config.Exporter) (*TelegramExporter, error) {
	opts, err := NewTelegramOptions(cfg.Options)
	if err != nil {
		return nil, err
	}
	bot, err := tgbotapi.NewBotAPI(opts.Token)
	if err != nil {
		return nil, err
	}
	return &TelegramExporter{
		Bot:     bot,
		ChatIDs: opts.ChatIDs,
	}, nil
}

func (te *TelegramExporter) SendMessage(msg string) error {
	for _, chatID := range te.ChatIDs {
		m := tgbotapi.NewMessage(int64(chatID), msg)
		_, err := te.Bot.Send(m)
		if err != nil {
			return err
		}
	}
	return nil
}
