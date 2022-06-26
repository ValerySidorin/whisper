//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/infrastructure/appctx"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web/routes"
	"github.com/google/wire"
)

func InitWebServer() (*web.Server, error) {
	wire.Build(
		//wire.Bind(new(port.ContextProvider), new(*appctx.CoreContext)),

		routes.Register,
		web.Register,
		appctx.Register,
		config.Register,
	)
	return &web.Server{}, nil
}
