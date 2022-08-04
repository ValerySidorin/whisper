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
		return nil, errors.New("token is not present")
	}
	t, ok := token.(string)
	if !ok {
		return nil, errors.New("token string type assertion failed")
	}
	return &telegramOptions{
		token: t,
	}, nil
}
