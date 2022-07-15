package gitlab

import (
	"errors"
)

type gitlabOptions struct {
	token string
	url   string
}

func NewGitlabOptions(opts map[string]interface{}) (*gitlabOptions, error) {
	token, ok := opts["token"]
	if !ok {
		return nil, errors.New("token is not present")
	}
	url := opts["url"]
	return &gitlabOptions{
		token: token.(string),
		url:   url.(string),
	}, nil
}
