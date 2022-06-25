package handler

import (
	"fmt"
	"net/http"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/messageBuilders"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Provider  string
	Exporters []config.Exporter
	Templates map[string]string
}

func New(cfg config.Handler) *Handler {
	return &Handler{
		Provider:  cfg.Provider,
		Exporters: cfg.Exporters,
		Templates: cfg.Templates,
	}
}

func (h *Handler) DefaultHandlerFunc(ctx *fasthttp.RequestCtx) {
	mb, err := messageBuilders.Get(h.Provider)
	if err != nil {
		h.processError(ctx, err)
		return
	}
	m, err := mb.Build(ctx.Request.Body(), h.Templates)
	if err != nil {
		h.processError(ctx, err)
		return
	}
	ctx.Response.SetStatusCode(http.StatusOK)
	ctx.Response.SetBodyString(m)
}

func (h *Handler) processError(ctx *fasthttp.RequestCtx, err error) {
	fmt.Println(err)
	apiErr, ok := err.(port.HttpError)
	if ok {
		code, msg := apiErr.GetCode(), apiErr.Error()
		ctx.SetStatusCode(code)
		if code >= http.StatusInternalServerError {
			ctx.Response.SetBodyString("Internal server error")
		} else {
			ctx.SetBodyString(msg)
		}
	} else {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBodyString("Internal server error")
	}
}
