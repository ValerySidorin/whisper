//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/appctx"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web/routes"
	"github.com/google/wire"
)

func InitWebServer() (*web.Server, error) {
	wire.Build(
		wire.Bind(new(port.MessageRenderer), new(*domain.DefaultMessageRenderer)),

		routes.Register,
		web.Register,
		appctx.Register,
		config.Register,
		messenger.Register,
		vcshosting.RegisterEventParser,
		domain.RegisterDefaultMessageRenderer,
	)
	return &web.Server{}, nil
}
