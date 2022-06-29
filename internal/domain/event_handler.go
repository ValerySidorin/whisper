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
	for _, v := range h.exporters {
		if err := v.SendMergeRequest(e); err != nil {
			return err
		}
	}
	return nil
}

func (h *EventHandler) HandleDeployment(body []byte) error {
	e, err := h.eventParser.ParseDeployment(body)
	if err != nil {
		return err
	}
	for _, v := range h.exporters {
		if err := v.SendDeployment(e); err != nil {
			return err
		}
	}
	return nil
}
