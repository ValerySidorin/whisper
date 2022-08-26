package domain

import (
	"reflect"
	"strconv"

	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/dto/storage"
	vcsdto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	"github.com/ValerySidorin/whisper/internal/domain/port"
)

type DefaultMessengerBot struct {
	storage   port.Storager
	renderer  port.MessageRenderer
	messenger port.MessengerService
}

func NewMessengerBot(s port.Storager, r port.MessageRenderer, m port.MessengerService) *DefaultMessengerBot {
	return &DefaultMessengerBot{
		storage:   s,
		renderer:  r,
		messenger: m,
	}
}

func (bot *DefaultMessengerBot) SendMergeRequestEvent(mre *vcsdto.MergeRequestEvent, chatID int64) error {
	msg := bot.renderer.RenderMergeRequestEvent(mre)
	if msg != "" {
		if err := bot.messenger.SendMessage(chatID, msg); err != nil {
			return nil
		}
	}
	return nil
}

func (bot *DefaultMessengerBot) SendDeploymentEvent(de *vcsdto.DeploymentEvent, chatID int64) error {
	msg := bot.renderer.RenderDeploymentEvent(de)
	if msg != "" {
		if err := bot.messenger.SendMessage(chatID, msg); err != nil {
			return nil
		}
	}
	return nil
}

func (bot *DefaultMessengerBot) SendBuildEvent(be *vcsdto.BuildEvent, chatID int64) error {
	msg := bot.renderer.RenderBuildEvent(be)
	if msg != "" {
		if err := bot.messenger.SendMessage(chatID, msg); err != nil {
			return nil
		}
	}
	return nil
}

func (bot *DefaultMessengerBot) HandleMessage(msg string, userID int64, vcs dto.VCSHostingType) {
	m := bot.messenger.GetMessengerType()
	if reflect.TypeOf(msg).Kind() == reflect.String && msg != "" {
		switch msg {
		case "/start":
			user, _ := bot.storage.GetUserByMessenger(vcs, m, userID)
			if user != nil {
				bot.messenger.SendMessage(userID, "Hey! You are registered already!")
			} else {
				user = &storage.User{
					VCSHostingType:  vcs,
					MessengerType:   dto.Telegram,
					State:           dto.Registering,
					MessengerUserID: userID,
				}
				_, err := bot.storage.AddUser(user)
				if err != nil {
					bot.messenger.SendMessage(userID, "An error occured while registering you in whisper app. Try again later.")
				} else {
					bot.messenger.SendMessage(userID, "Hello! Let's try to register! Give me your vcs user ID:")
				}
			}
		default:
			user, err := bot.storage.GetUserByMessenger(vcs, m, userID)
			if err != nil {
				bot.messenger.SendMessage(userID, "I don't know you! Type /start to register in whisper app.")
			} else {
				if user.State == dto.Registering {
					vcsID, err := strconv.Atoi(msg)
					if err != nil {
						bot.messenger.SendMessage(userID, "ID is a number, you stupid! Still waiting...")
					} else {
						user.VCSHostingUserID = int64(vcsID)
						user.State = dto.Idle
						_, err := bot.storage.UpdateUser(user)
						if err != nil {
							bot.messenger.SendMessage(userID, "An error occured while saving your data to my DB! Try again later!")
						} else {
							bot.messenger.SendMessage(userID, "Gracias! You are registered now!")
						}
					}
				} else {
					bot.messenger.SendMessage(userID, "I don't know this command...")
				}
			}
		}
	}
}
