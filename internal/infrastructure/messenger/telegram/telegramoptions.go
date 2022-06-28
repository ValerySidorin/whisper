package telegram

import (
	"errors"
)

type TelegramOptions struct {
	Token   string
	ChatIDs []int
}

func NewTelegramOptions(opts map[string]interface{}) (*TelegramOptions, error) {
	token, ok := opts["token"]
	if !ok {
		return nil, errors.New("telegram token is not present")
	}
	chatIDs, ok := opts["chatIds"]
	if !ok {
		return nil, errors.New("telegram chat IDs not present")
	}
	intIDs := make([]int, 0)
	for _, v := range chatIDs.([]interface{}) {
		intIDs = append(intIDs, v.(int))
	}
	return &TelegramOptions{
		Token:   token.(string),
		ChatIDs: intIDs,
	}, nil
}
