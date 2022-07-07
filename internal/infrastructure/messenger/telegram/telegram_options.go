package telegram

import (
	"errors"
)

type TelegramOptions struct {
	Token string
}

func NewTelegramOptions(opts map[string]interface{}) (*TelegramOptions, error) {
	token, ok := opts["token"]
	if !ok {
		return nil, errors.New("telegram token is not present")
	}
	return &TelegramOptions{
		Token: token.(string),
	}, nil
}
