package exporters

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramExporter struct {
	Bot *tgbotapi.BotAPI
}

func NewTelegramExporter(token string) (*TelegramExporter, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &TelegramExporter{
		Bot: bot,
	}, nil
}

func (te *TelegramExporter) SendMessage(msg string, chatID int64) error {
	m := tgbotapi.NewMessage(chatID, msg)
	_, err := te.Bot.Send(m)
	if err != nil {
		return err
	}
	return nil
}
