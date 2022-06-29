package gitlab

import (
	"encoding/json"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/converters"
	gitlab "github.com/xanzy/go-gitlab"
)

type GitlabHandler struct {
	Exporters []port.Exporter
	Client    *gitlab.Client
}

func NewHandler(cfg *config.Handler) (*GitlabHandler, error) {
	opts, err := NewGitlabOptions(cfg.VCSHosting.Options)
	if err != nil {
		return nil, err
	}
	c, err := gitlab.NewClient(opts.Token, gitlab.WithBaseURL(opts.URL))
	exporters := make([]port.Exporter, 0)
	for _, v := range cfg.Exporters {
		e, err := messenger.GetExporter(&v)
		if err != nil {
			return nil, err
		}
		exporters = append(exporters, e)
	}
	if err != nil {
		return nil, err
	}
	return &GitlabHandler{
		Client:    c,
		Exporters: exporters,
	}, nil
}

func (gh *GitlabHandler) HandleMergeRequest(body []byte) error {
	gmr := gitlab.MergeEvent{}
	if err := json.Unmarshal(body, &gmr); err != nil {
		return err
	}
	conv := converters.MRConverter{MergeEvent: &gmr}
	m, err := conv.Convert()
	if err != nil {
		return err
	}
	msg := m.GetMessage()
	if msg != "" {
		for _, e := range gh.Exporters {
			if err := e.SendMessage(msg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (gh *GitlabHandler) HandleDeployment(body []byte) error {
	gd := gitlab.DeploymentEvent{}
	if err := json.Unmarshal(body, &gd); err != nil {
		return err
	}
	j, _, err := gh.Client.Jobs.GetJob(gd.Project.ID, gd.DeployableID)
	if err != nil {
		return err
	}
	conv := converters.DeploymentConverter{DeploymentEvent: &gd, Job: j}
	m, err := conv.Convert()
	if err != nil {
		return err
	}
	msg := m.GetMessage()
	if msg != "" {
		for _, e := range gh.Exporters {
			if err := e.SendMessage(msg); err != nil {
				return err
			}
		}
	}
	return nil
}
