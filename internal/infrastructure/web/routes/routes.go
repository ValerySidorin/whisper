package routes

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web/handler"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Register(cfg config.Configuration) (*router.Router, error) {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBodyString("Whisper is your CI event notifier")
	})
	for _, v := range cfg.Handlers {
		h, err := handler.New(v)
		if err != nil {
			return nil, err
		}
		r.POST(v.Route, h.DefaultHandlerFunc)
	}
	return r, nil
}
