package gitlab

import (
	"errors"
)

type GitlabOptions struct {
	Token string
	URL   string
}

func NewGitlabOptions(opts map[string]interface{}) (*GitlabOptions, error) {
	token, ok := opts["token"]
	if !ok {
		return nil, errors.New("gitlab token is not present")
	}
	url := opts["url"]
	return &GitlabOptions{
		Token: token.(string),
		URL:   url.(string),
	}, nil
}
