package telegram

import (
	"fmt"

	"github.com/ValerySidorin/whisper/internal/domain"
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramMessenger struct {
	baseBot port.MessengerBot
	bot     *tgbotapi.BotAPI
}

func Register(cfg *config.Configuration, r port.MessageRenderer, storage port.Storager) (*TelegramMessenger, error) {
	opts, err := newTelegramOptions(cfg.Messenger.Options)
	if err != nil {
		return nil, fmt.Errorf("telegram: malformed options: %s", err)
	}
	bot, err := tgbotapi.NewBotAPI(opts.token)
	if err != nil {
		return nil, fmt.Errorf("telegram: failed to create bot: %s", err)
	}
	m := &TelegramMessenger{
		bot: bot,
	}
	m.baseBot = domain.NewMessengerBot(storage, r, m)
	go m.telegramBotListenAndServe(cfg.VCSHosting.Provider)
	return m, nil
}

func (te *TelegramMessenger) SendMessage(chatID int64, msg string) error {
	m := tgbotapi.NewMessage(int64(chatID), msg)
	_, err := te.bot.Send(m)
	if err != nil {
		return fmt.Errorf("telegram: failed send message: %s", err)
	}
	return nil
}

func (te *TelegramMessenger) GetMessengerType() dto.MessengerType {
	return dto.Telegram
}

func (te *TelegramMessenger) telegramBotListenAndServe(vcs dto.VCSHostingType) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := te.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		te.baseBot.HandleMessage(update.Message.Text, update.SentFrom().ID, vcs)
	}
}
