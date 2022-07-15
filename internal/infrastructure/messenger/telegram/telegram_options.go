package telegram

import (
	"errors"
)

type telegramOptions struct {
	token string
}

func newTelegramOptions(opts map[string]interface{}) (*telegramOptions, error) {
	token, ok := opts["token"]
	if !ok {
		return nil, errors.New("telegram: token is not present")
	}
	return &telegramOptions{
		token: token.(string),
	}, nil
}
