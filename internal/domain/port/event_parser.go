package port

import "github.com/ValerySidorin/whisper/internal/domain/dto"

type EventParser interface {
	ParseMergeRequest(body []byte) (*dto.MergeRequest, error)
	ParseDeployment(body []byte) (*dto.Deployment, error)
}
