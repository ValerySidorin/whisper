package port

import "github.com/ValerySidorin/whisper/internal/domain/dto"

type MessengerService interface {
	SendMessage(chatID int64, msg string) error
	GetMessengerType() dto.MessengerType
}
