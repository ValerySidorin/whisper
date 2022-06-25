package messageBuilders

import "fmt"

type MessageBuilder interface {
	Build(event []byte, templates map[string]string) (string, error)
}

func Get(eventProvider string) (MessageBuilder, error) {
	switch eventProvider {
	case "gitlab":
		return &GitlabMessageBuilder{}, nil
	default:
		return nil, fmt.Errorf("unknown event provider: %v", eventProvider)
	}

}
