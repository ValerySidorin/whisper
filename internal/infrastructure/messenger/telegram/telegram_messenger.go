package telegram

import (
	"reflect"
	"strconv"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/dto/storage"
	vcsdto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramMessenger struct {
	cfg      *config.Configuration
	storage  port.Storager
	renderer port.MessageRenderer
	bot      *tgbotapi.BotAPI
	chatIDs  []int
}

func Register(cfg *config.Configuration, r port.MessageRenderer, storage port.Storager) (*TelegramMessenger, error) {
	opts, err := NewTelegramOptions(cfg.Messenger.Options)
	if err != nil {
		return nil, err
	}
	bot, err := tgbotapi.NewBotAPI(opts.Token)
	if err != nil {
		return nil, err
	}
	m := &TelegramMessenger{
		cfg:      cfg,
		storage:  storage,
		bot:      bot,
		chatIDs:  opts.ChatIDs,
		renderer: r,
	}
	go m.BotListenAndServe()
	return m, nil
}

func (te *TelegramMessenger) SendMergeRequest(mr *vcsdto.MergeRequest) error {
	msg := te.renderer.RenderMergeRequest(mr)
	if msg != "" {
		for _, chatID := range te.chatIDs {
			if err := te.SendMessage(int64(chatID), msg); err != nil {
				return nil
			}
		}
	}
	return nil
}

func (te *TelegramMessenger) SendDeployment(d *vcsdto.Deployment) error {
	msg := te.renderer.RenderDeployment(d)
	if msg != "" {
		for _, chatID := range te.chatIDs {
			if err := te.SendMessage(int64(chatID), msg); err != nil {
				return nil
			}
		}
	}
	return nil
}

func (te *TelegramMessenger) SendMessage(chatID int64, msg string) error {
	m := tgbotapi.NewMessage(int64(chatID), msg)
	_, err := te.bot.Send(m)
	if err != nil {
		return err
	}
	return nil
}

func (te *TelegramMessenger) BotListenAndServe() {
	vcsType := te.cfg.VCSHosting.Provider
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := te.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				user, err := te.storage.GetUser(vcsType, dto.Telegram, update.SentFrom().ID)
				if err != nil {
					te.SendMessage(int64(update.SentFrom().ID), "An error occured while fetching your data from my DB! Try again later!")
				} else {
					if user != nil {
						te.SendMessage(int64(update.SentFrom().ID), "Hey! You are registered already!")
					} else {
						user = &storage.User{
							VCSHostingType:  vcsType,
							MessengerType:   dto.Telegram,
							State:           dto.Registering,
							MessengerUserID: update.SentFrom().ID,
						}
						_, err := te.storage.AddUser(user)
						if err != nil {
							te.SendMessage(int64(update.SentFrom().ID), "An error occured while registering you in whisper app. Try again later.")
						} else {
							te.SendMessage(int64(update.SentFrom().ID), "Hello! Let's try to register! Give me your vcs user ID:")
						}
					}
				}

			default:
				user, err := te.storage.GetUser(vcsType, dto.Telegram, update.SentFrom().ID)
				if err != nil {
					te.SendMessage(int64(update.SentFrom().ID), "I don't know you! Type /start to register in whisper app.")
				} else {
					if user.State == dto.Registering {
						vcsID, err := strconv.Atoi(update.Message.Text)
						if err != nil {
							te.SendMessage(int64(update.SentFrom().ID), "ID is a number, you stupid! Still waiting...")
						} else {
							user.VCSHostingUserID = int64(vcsID)
							user.State = dto.Idle
							_, err := te.storage.UpdateUser(user)
							if err != nil {
								te.SendMessage(int64(update.SentFrom().ID), "An error occured while saving your data to my DB! Try again later!")
							} else {
								te.SendMessage(int64(update.SentFrom().ID), "Gracias! You are registered now!")
							}
						}
					} else {
						te.SendMessage(int64(update.SentFrom().ID), "I don't know this command...")
					}
				}
			}
		}
	}
}
