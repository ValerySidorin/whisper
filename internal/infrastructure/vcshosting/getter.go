package vcshosting

import (
	"errors"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab"
)

func GetEventParser(cfg *config.Handler) (port.EventParser, error) {
	switch cfg.VCSHosting.Provider {
	case "gitlab":
		return gitlab.NewEventParser(cfg)
	default:
		return nil, errors.New("unknown vcshosting")
	}
}
