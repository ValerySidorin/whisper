package port

import "github.com/ValerySidorin/whisper/internal/domain/dto"

type MessageRenderer interface {
	RenderMergeRequest(*dto.MergeRequest) string
	RenderDeployment(*dto.Deployment) string
}
