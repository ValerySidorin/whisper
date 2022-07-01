package port

import dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"

type EventParser interface {
	ParseMergeRequest(body []byte) (*dto.MergeRequest, error)
	ParseDeployment(body []byte) (*dto.Deployment, error)
}
