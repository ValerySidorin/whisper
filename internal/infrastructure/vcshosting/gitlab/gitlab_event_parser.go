package gitlab

import (
	"encoding/json"
	"fmt"

	"github.com/ValerySidorin/whisper/internal/config"
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/converters"
	"github.com/xanzy/go-gitlab"
)

type GitlabEventParser struct {
	Client *gitlab.Client
}

func RegisterEventParser(cfg *config.Configuration) (*GitlabEventParser, error) {
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

func (p *GitlabEventParser) ParseMergeRequestEvent(body []byte) (*dto.MergeRequestEvent, error) {
	gmr := gitlab.MergeEvent{}
	if err := json.Unmarshal(body, &gmr); err != nil {
		return nil, err
	}
	a, _, err := p.Client.Users.GetUser(gmr.ObjectAttributes.AuthorID, gitlab.GetUsersOptions{})
	if err != nil {
		return nil, fmt.Errorf("error fetching merge request author: %s", err)
	}
	conv := converters.MREventConverter{MergeEvent: &gmr, User: a}
	m, err := conv.Convert()
	if err != nil {
		return nil, err
	}
	return m.(*dto.MergeRequestEvent), nil
}

func (p *GitlabEventParser) ParseDeploymentEvent(body []byte) (*dto.DeploymentEvent, error) {
	gd := gitlab.DeploymentEvent{}
	if err := json.Unmarshal(body, &gd); err != nil {
		return nil, err
	}
	j, _, err := p.Client.Jobs.GetJob(gd.Project.ID, gd.DeployableID)
	if err != nil {
		return nil, err
	}
	conv := converters.DeploymentEventConverter{DeploymentEvent: &gd, Job: j}
	m, err := conv.Convert()
	if err != nil {
		return nil, err
	}
	return m.(*dto.DeploymentEvent), nil
}
