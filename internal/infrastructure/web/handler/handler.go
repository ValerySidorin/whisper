package handler

import (
	"fmt"
	"net/http"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/exporters"
	"github.com/ValerySidorin/whisper/internal/domain/messagebuilders"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Provider  string
	Exporters []config.Exporter
	Templates map[string]string
}

func New(cfg config.Handler) (*Handler, error) {
	return &Handler{
		Provider:  cfg.Provider,
		Exporters: cfg.Exporters,
		Templates: cfg.Templates,
	}, nil
}

func (h *Handler) DefaultHandlerFunc(ctx *fasthttp.RequestCtx) {
	mb, err := messagebuilders.Get(h.Provider)
	if err != nil {
		h.processError(ctx, err)
		return
	}
	m, err := mb.Build(ctx.Request.Body(), h.Templates)
	if err != nil {
		h.processError(ctx, err)
		return
	}
	for _, v := range h.Exporters {
		e, err := exporters.Get(&v)
		if err != nil {
			h.processError(ctx, err)
			return
		}
		for _, chatID := range v.ChatIds {
			if err := e.SendMessage(m, chatID); err != nil {
				h.processError(ctx, err)
				return
			}
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
