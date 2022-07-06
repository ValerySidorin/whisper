package gitlab

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	"github.com/xanzy/go-gitlab"
)

type GitlabService struct {
	Client *gitlab.Client
}

func Register(cfg *config.Configuration) (*GitlabService, error) {
	opts, err := NewGitlabOptions(cfg.VCSHosting.Options)
	if err != nil {
		return nil, err
	}
	c, err := gitlab.NewClient(opts.Token, gitlab.WithBaseURL(opts.URL))
	if err != nil {
		return nil, err
	}
	return &GitlabService{Client: c}, nil
}

func (gs *GitlabService) GetPerson(username string) (*vcshosting.Person, error) {
	users, _, err := gs.Client.Search.Users("username="+username, nil)
	if err != nil {
		return nil, err
	}
	u := users[0]
	return &vcshosting.Person{
		Name:     u.Name,
		UserName: u.Username,
	}, nil
}
