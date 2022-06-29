package port

import "github.com/ValerySidorin/whisper/internal/domain/dto"

type Exporter interface {
	SendMergeRequest(mr *dto.MergeRequest) error
	SendDeployment(d *dto.Deployment) error
}
