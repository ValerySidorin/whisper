package routes

import (
	"github.com/ValerySidorin/whisper/internal/config"
	handler "github.com/ValerySidorin/whisper/internal/infrastructure/web/handlers"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Register(cfg config.Configuration) *router.Router {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBodyString("Whisper is your CI event notifier")
	})
	for _, v := range cfg.Handlers {
		h := handler.New(v)
		r.POST(v.Route, h.DefaultHandlerFunc)
	}
	return r
}
