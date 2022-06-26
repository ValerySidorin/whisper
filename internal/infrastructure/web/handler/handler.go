package handler

import (
	"fmt"
	"net/http"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Provider  string
	Exporters []config.Exporter
}

func New(cfg config.Handler) (*Handler, error) {
	return &Handler{
		Provider:  cfg.Provider,
		Exporters: cfg.Exporters,
	}, nil
}

func (h *Handler) DefaultHandlerFunc(ctx *fasthttp.RequestCtx) {
	m, err := vcshosting.GetMessageable(h.Provider, string(ctx.Request.Body()))
	if err != nil {
		h.processError(ctx, err)
		return
	}
	for _, v := range h.Exporters {
		e, err := messenger.GetExporter(&v)
		if err != nil {
			h.processError(ctx, err)
			return
		}
		if err := e.SendMessage(m.GetMessage()); err != nil {
			h.processError(ctx, err)
			return
		}
	}
	ctx.Response.SetStatusCode(http.StatusOK)
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
