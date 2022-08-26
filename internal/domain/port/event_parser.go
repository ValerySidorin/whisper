package port

import dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"

type EventParser interface {
	ParseMergeRequestEvent(body []byte) (*dto.MergeRequestEvent, error)
	ParseDeploymentEvent(body []byte) (*dto.DeploymentEvent, error)
	ParseBuildEvent(body []byte) (*dto.BuildEvent, error)
}
