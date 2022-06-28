package gitlab

import (
	"encoding/json"

	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/converters"
	gitlabdto "github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/dto"
)

type GitlabHandler struct {
	Exporters []port.Exporter
}

func (gh *GitlabHandler) HandleMergeRequest(body []byte) (*dto.MergeRequest, error) {
	gmr := gitlabdto.MergeRequest{}
	if err := json.Unmarshal(body, &gmr); err != nil {
		return nil, err
	}
	converter := converters.MRConverter{MergeRequest: &gmr}
	m, err := converter.Convert()
	mr := m.(*dto.MergeRequest)
	if err != nil {
		return nil, err
	}
	for _, e := range gh.Exporters {
		if err := e.SendMessage(mr.GetMessage()); err != nil {
			return nil, err
		}
	}
	return mr, nil
}

func (gh *GitlabHandler) HandleDeployment(body []byte) (*dto.Deployment, error) {
	return nil, nil
}
