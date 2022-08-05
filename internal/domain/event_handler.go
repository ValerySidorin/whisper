package domain

import (
	"fmt"

	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/port"
)

type EventHandler struct {
	vcsType        dto.VCSHostingType
	messengerType  dto.MessengerType
	baseBot        port.MessengerBot
	eventParser    port.EventParser
	storage        port.Storager
	defaultChatIDs []int64
}

func NewEventHandler(
	vscType dto.VCSHostingType,
	messengerType dto.MessengerType,
	b port.MessengerBot,
	p port.EventParser,
	s port.Storager,
	cIDs []int64) *EventHandler {
	return &EventHandler{
		vcsType:        vscType,
		messengerType:  messengerType,
		baseBot:        b,
		eventParser:    p,
		storage:        s,
		defaultChatIDs: cIDs,
	}
}

func (h *EventHandler) HandleMergeRequest(body []byte) error {
	e, err := h.eventParser.ParseMergeRequestEvent(body)
	if err != nil {
		return fmt.Errorf("domain: error parsing merge request: %s", err)
	}
	assignee, _ := h.storage.GetUserByVCSHosting(h.vcsType, h.messengerType, e.MergeRequest.Assignee.ID)
	author, _ := h.storage.GetUserByVCSHosting(h.vcsType, h.messengerType, e.MergeRequest.Author.ID)
	if e.MergeRequest.State == "opened" {
		if assignee != nil {
			if err := h.baseBot.SendMergeRequestEvent(e, assignee.MessengerUserID); err != nil {
				return fmt.Errorf("domain: error sending merge request to assignee: %s", err)
			}
		}
	}
	if e.MergeRequest.State == "merged" || e.MergeRequest.State == "closed" {
		if author != nil {
			if err := h.baseBot.SendMergeRequestEvent(e, author.MessengerUserID); err != nil {
				return fmt.Errorf("domain: error sending merge request to author: %s", err)
			}
		}
	}
	for _, v := range h.defaultChatIDs {
		if err := h.baseBot.SendMergeRequestEvent(e, v); err != nil {
			return fmt.Errorf("domain: error sending merge request to chat %v: %s", v, err)
		}
	}
	return nil
}

func (h *EventHandler) HandleDeployment(body []byte) error {
	e, err := h.eventParser.ParseDeploymentEvent(body)
	if err != nil {
		return fmt.Errorf("domain: error parsing deployment: %s", err)
	}
	for _, v := range h.defaultChatIDs {
		if err := h.baseBot.SendDeploymentEvent(e, v); err != nil {
			return fmt.Errorf("domain: error sending deployment to chat %v: %s", v, err)
		}
	}
	return nil
}
