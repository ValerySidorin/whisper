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
	t, ok := token.(string)
	if !ok {
		return nil, errors.New("token string type assertion failed")
	}
	u, ok := url.(string)
	if !ok {
		return nil, errors.New("url string type assertion failed")
	}
	return &gitlabOptions{
		token: t,
		url:   u,
	}, nil
}
