package telegram

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramExporter struct {
	renderer port.MessageRenderer
	bot      *tgbotapi.BotAPI
	chatIDs  []int
}

func NewExporter(cfg *config.Exporter, r port.MessageRenderer) (*TelegramExporter, error) {
	opts, err := NewTelegramOptions(cfg.Options)
	if err != nil {
		return nil, err
	}
	bot, err := tgbotapi.NewBotAPI(opts.Token)
	if err != nil {
		return nil, err
	}
	return &TelegramExporter{
		bot:      bot,
		chatIDs:  opts.ChatIDs,
		renderer: r,
	}, nil
}

func (te *TelegramExporter) SendMergeRequest(mr *dto.MergeRequest) error {
	msg := te.renderer.RenderMergeRequest(mr)
	if msg != "" {
		for _, chatID := range te.chatIDs {
			m := tgbotapi.NewMessage(int64(chatID), msg)
			_, err := te.bot.Send(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (te *TelegramExporter) SendDeployment(d *dto.Deployment) error {
	msg := te.renderer.RenderDeployment(d)
	if msg != "" {
		for _, chatID := range te.chatIDs {
			m := tgbotapi.NewMessage(int64(chatID), msg)
			_, err := te.bot.Send(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
