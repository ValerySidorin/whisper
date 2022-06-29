package handler

import (
	"fmt"
	"net/http"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	eventHandler *domain.EventHandler
}

func New(cfg *config.Handler, r port.MessageRenderer) (*Handler, error) {
	exporters := make([]port.Exporter, 0)
	for _, v := range cfg.Exporters {
		e, err := messenger.GetExporter(&v, r)
		if err != nil {
			return nil, err
		}
		exporters = append(exporters, e)
	}
	p, err := vcshosting.GetEventParser(cfg)
	if err != nil {
		return nil, err
	}
	eh := domain.NewEventHandler(exporters, p)
	h := &Handler{}
	h.eventHandler = eh
	return h, nil
}

func (h *Handler) MergeRequestHandlerFunc(ctx *fasthttp.RequestCtx) {
	if err := h.eventHandler.HandleMergeRequest(ctx.Request.Body()); err != nil {
		h.processError(ctx, err)
	}
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (h *Handler) DeploymentHandlerFunc(ctx *fasthttp.RequestCtx) {
	if err := h.eventHandler.HandleDeployment(ctx.Request.Body()); err != nil {
		h.processError(ctx, err)
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
			ctx.Response.SetBodyString(apiErr.Error())
		} else {
			ctx.SetBodyString(msg)
		}
	} else {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBodyString(apiErr.Error())
	}
}
