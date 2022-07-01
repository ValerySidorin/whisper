package port

import dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"

type MessageRenderer interface {
	RenderMergeRequest(*dto.MergeRequest) string
	RenderDeployment(*dto.Deployment) string
}
