package port

import dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"

type Messenger interface {
	SendMergeRequest(mr *dto.MergeRequest) error
	SendDeployment(d *dto.Deployment) error
	SendMessage(chatID int64, msg string) error
	BotListenAndServe()
}
