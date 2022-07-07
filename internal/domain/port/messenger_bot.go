package port

import (
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	vcsdto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
)

type MessengerBot interface {
	SendMergeRequestEvent(mre *vcsdto.MergeRequestEvent, chatID int64) error
	SendDeploymentEvent(de *vcsdto.DeploymentEvent, chatID int64) error
	HandleMessage(msg string, userID int64, vcs dto.VCSHostingType)
}
