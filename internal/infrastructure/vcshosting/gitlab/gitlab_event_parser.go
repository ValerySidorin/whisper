package gitlab

import (
	"encoding/json"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/converters"
	"github.com/xanzy/go-gitlab"
)

type GitlabEventParser struct {
	Client *gitlab.Client
}

func NewEventParser(cfg *config.Handler) (*GitlabEventParser, error) {
	opts, err := NewGitlabOptions(cfg.VCSHosting.Options)
	if err != nil {
		return nil, err
	}
	c, err := gitlab.NewClient(opts.Token, gitlab.WithBaseURL(opts.URL))
	if err != nil {
		return nil, err
	}
	return &GitlabEventParser{Client: c}, nil
}

func (p *GitlabEventParser) ParseMergeRequest(body []byte) (*dto.MergeRequest, error) {
	gmr := gitlab.MergeEvent{}
	if err := json.Unmarshal(body, &gmr); err != nil {
		return nil, err
	}
	conv := converters.MRConverter{MergeEvent: &gmr}
	m, err := conv.Convert()
	if err != nil {
		return nil, err
	}
	return m.(*dto.MergeRequest), nil
}

func (p *GitlabEventParser) ParseDeployment(body []byte) (*dto.Deployment, error) {
	gd := gitlab.DeploymentEvent{}
	if err := json.Unmarshal(body, &gd); err != nil {
		return nil, err
	}
	j, _, err := p.Client.Jobs.GetJob(gd.Project.ID, gd.DeployableID)
	if err != nil {
		return nil, err
	}
	conv := converters.DeploymentConverter{DeploymentEvent: &gd, Job: j}
	m, err := conv.Convert()
	if err != nil {
		return nil, err
	}
	return m.(*dto.Deployment), nil
}
