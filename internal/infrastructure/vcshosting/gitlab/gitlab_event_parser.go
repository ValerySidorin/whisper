package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"

	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	"github.com/ValerySidorin/whisper/internal/infrastructure/config"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/converters"
	"github.com/xanzy/go-gitlab"
)

type GitlabEventParser struct {
	client *gitlab.Client
}

func RegisterEventParser(cfg *config.Configuration) (*GitlabEventParser, error) {
	opts, err := NewGitlabOptions(cfg.VCSHosting.Options)
	if err != nil {
		return nil, fmt.Errorf("gitlab: malformed gitlab options: %s", err)
	}
	c, err := gitlab.NewClient(opts.token, gitlab.WithBaseURL(opts.url))
	if err != nil {
		return nil, fmt.Errorf("gitlab: failed to create client: %s", err)
	}
	return &GitlabEventParser{client: c}, nil
}

func (p *GitlabEventParser) ParseMergeRequestEvent(body []byte) (*dto.MergeRequestEvent, error) {
	gmr := gitlab.MergeEvent{}
	if err := json.Unmarshal(body, &gmr); err != nil {
		return nil, errors.New("gitlab: merge request event unmarshalling error")
	}
	a, _, err := p.client.Users.GetUser(gmr.ObjectAttributes.AuthorID, gitlab.GetUsersOptions{})
	if err != nil {
		return nil, fmt.Errorf("gitlab: fetching user error: %s", err)
	}
	conv := converters.MREventConverter{MergeEvent: &gmr, User: a}
	m := conv.Convert()
	return m.(*dto.MergeRequestEvent), nil
}

func (p *GitlabEventParser) ParseDeploymentEvent(body []byte) (*dto.DeploymentEvent, error) {
	gd := gitlab.DeploymentEvent{}
	if err := json.Unmarshal(body, &gd); err != nil {
		return nil, errors.New("gitlab: deployment event unmarshalling error")
	}
	j, _, err := p.client.Jobs.GetJob(gd.Project.ID, gd.DeployableID)
	if err != nil {
		return nil, fmt.Errorf("gitlab: get job error: %s", err)
	}
	conv := converters.DeploymentEventConverter{DeploymentEvent: &gd, Job: j}
	m := conv.Convert()
	return m.(*dto.DeploymentEvent), nil
}
