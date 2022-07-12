package routes

import (
	"github.com/ValerySidorin/whisper/internal/infrastructure/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web/handler"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Register(cfg *config.Configuration, b port.MessengerBot, p port.EventParser, rndr port.MessageRenderer, s port.Storager) (*router.Router, error) {
	r := router.New()
	r.GET("/", DefaultHandlerFunc)
	for _, v := range cfg.Handlers {
		h, err := handler.New(cfg, b, p, s, v.DefaultChatIDs)
		if err != nil {
			return nil, err
		}
		switch v.Action {
		case "merge_request":
			r.POST(v.Route, h.MergeRequestHandlerFunc)
		case "deployment":
			r.POST(v.Route, h.DeploymentHandlerFunc)
		default:
			r.POST(v.Route, DefaultHandlerFunc)
		}
	}
	return r, nil
}

func DefaultHandlerFunc(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetBodyString("Whisper is your CI event notifier")
}
