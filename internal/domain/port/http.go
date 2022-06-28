package port

import "github.com/ValerySidorin/whisper/internal/domain/dto"

type HttpError interface {
	error
	GetCode() int
	GetErr() error
}

type VCSHostingHandler interface {
	HandleMergeRequest(body []byte) (*dto.MergeRequest, error)
	HandleDeployment(body []byte) (*dto.Deployment, error)
}
