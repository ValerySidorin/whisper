package port

import (
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/dto/storage"
)

type Storager interface {
	AddUser(u *storage.User) (*storage.User, error)
	GetUser(vcsType dto.VCSHostingType, messengerType dto.MessengerType, messengerUserID int64) (*storage.User, error)
	UpdateUser(u *storage.User) (*storage.User, error)
}
