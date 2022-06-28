package vcshosting

import (
	"errors"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab"
)

func GetVCSHostingHandler(cfg *config.Handler) (port.VCSHostingHandler, error) {
	switch cfg.VCSHosting.Provider {
	case "gitlab":
		return gitlab.NewHandler(cfg)
	default:
		return nil, errors.New("unknown vcshosting")
	}
}
