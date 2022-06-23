package routes

import (
	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/infrastructure/web/handlers"
	"github.com/fasthttp/router"
)

func ProvideRoutes(h *handlers.Handlers, cfg *config.Configuration) *router.Router {
	return getDomainRouter(h, &cfg.Routes)
}

func getDomainRouter(h *handlers.Handlers, cfg *config.Routes) *router.Router {
	r := router.New()
	for _, v := range cfg.GitlabRoutes {
		r.POST(v, h.DefaultHandler)
	}
	return r
}
