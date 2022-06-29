package domain

import (
	"github.com/ValerySidorin/whisper/internal/domain/port"
)

type EventHandler struct {
	exporters   []port.Exporter
	eventParser port.EventParser
}

func NewEventHandler(
	exps []port.Exporter,
	p port.EventParser) *EventHandler {
	return &EventHandler{
		exporters:   exps,
		eventParser: p,
	}
}

func (h *EventHandler) HandleMergeRequest(body []byte) error {
	e, err := h.eventParser.ParseMergeRequest(body)
	if err != nil {
		return err
	}
	return h.sendEvent(e)
}

func (h *EventHandler) HandleDeployment(body []byte) error {
	e, err := h.eventParser.ParseDeployment(body)
	if err != nil {
		return err
	}
	return h.sendEvent(e)
}

func (h *EventHandler) sendEvent(m port.Messageable) error {
	msg := m.GetMessage()
	if msg != "" {
		for _, v := range h.exporters {
			if err := v.SendMessage(msg); err != nil {
				return err
			}
		}
	}
	return nil
}
