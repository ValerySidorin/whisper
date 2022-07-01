package domain

import (
	"github.com/ValerySidorin/whisper/internal/domain/port"
)

type EventHandler struct {
	messenger   port.Messenger
	eventParser port.EventParser
}

func NewEventHandler(
	m port.Messenger,
	p port.EventParser) *EventHandler {
	return &EventHandler{
		messenger:   m,
		eventParser: p,
	}
}

func (h *EventHandler) HandleMergeRequest(body []byte) error {
	e, err := h.eventParser.ParseMergeRequest(body)
	if err != nil {
		return err
	}
	if err := h.messenger.SendMergeRequest(e); err != nil {
		return err
	}
	return nil
}

func (h *EventHandler) HandleDeployment(body []byte) error {
	e, err := h.eventParser.ParseDeployment(body)
	if err != nil {
		return err
	}
	if err := h.messenger.SendDeployment(e); err != nil {
		return err
	}
	return nil
}
