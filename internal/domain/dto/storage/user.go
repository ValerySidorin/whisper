package storage

import "github.com/ValerySidorin/whisper/internal/domain/dto"

type User struct {
	ID               int64              `json:"id"`
	VCSHostingType   dto.VCSHostingType `json:"vcs_hosting_type"`
	MessengerType    dto.MessengerType  `json:"messenger_type"`
	VCSHostingUserID int64              `json:"vcs_hosting_user_id"`
	MessengerUserID  int64              `json:"messenger_user_id"`
	State            dto.UserState      `json:"state"`
}
