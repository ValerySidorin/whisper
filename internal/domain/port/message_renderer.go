package port

import dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"

type MessageRenderer interface {
	RenderMergeRequestEvent(*dto.MergeRequestEvent) string
	RenderDeploymentEvent(*dto.DeploymentEvent) string
	RenderBuildEvent(*dto.BuildEvent) string
}
