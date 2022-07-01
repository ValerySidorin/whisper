package vcshosting

import (
	"errors"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab"
)

func RegisterEventParser(cfg *config.Configuration) (port.EventParser, error) {
	switch cfg.VCSHosting.Provider {
	case "gitlab":
		return gitlab.NewEventParser(&cfg.VCSHosting)
	default:
		return nil, errors.New("unknown vcshosting")
	}
}
