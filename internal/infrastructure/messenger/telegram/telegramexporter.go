package telegram

import (
	"errors"

	"github.com/ValerySidorin/whisper/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramExporter struct {
	Bot     *tgbotapi.BotAPI
	ChatIDs []int
}

func NewExporter(cfg *config.Exporter) (*TelegramExporter, error) {
	token, ok := cfg.Attributes["token"]
	if !ok {
		return nil, errors.New("telegram token is not present")
	}
	chatIDs, ok := cfg.Attributes["chatIds"]
	if !ok {
		return nil, errors.New("telegram chat IDs not present")
	}

	IDs := make([]int, 0)
	for _, v := range chatIDs.([]interface{}) {
		IDs = append(IDs, v.(int))
	}

	bot, err := tgbotapi.NewBotAPI(token.(string))
	if err != nil {
		return nil, err
	}
	return &TelegramExporter{
		Bot:     bot,
		ChatIDs: IDs,
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
