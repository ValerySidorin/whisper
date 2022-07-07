//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/appctx"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger/telegram"
	"github.com/ValerySidorin/whisper/internal/infrastructure/storage/gorm"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web/routes"
	"github.com/google/wire"
)

func InitWebServer() (*web.Server, error) {
	wire.Build(
		wire.Bind(new(port.MessageRenderer), new(*domain.DefaultMessageRenderer)),
		wire.Bind(new(port.Messenger), new(*telegram.TelegramMessenger)),
		wire.Bind(new(port.EventParser), new(*gitlab.GitlabEventParser)),
		wire.Bind(new(port.Storager), new(*gorm.GormStorage)),

		routes.Register,
		web.Register,
		appctx.Register,
		config.Register,
		telegram.Register,
		gitlab.RegisterEventParser,
		domain.RegisterDefaultMessageRenderer,
		gorm.Register,
	)
	return &web.Server{}, nil
}
