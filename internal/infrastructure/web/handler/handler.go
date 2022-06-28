package handler

import (
	"fmt"
	"net/http"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/messenger"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Handler port.VCSHostingHandler
}

func New(cfg *config.Handler) (*Handler, error) {
	exporters := make([]port.Exporter, 0)
	for _, v := range cfg.Exporters {
		e, err := messenger.GetExporter(&v)
		if err != nil {
			return nil, err
		}
		exporters = append(exporters, e)
	}
	h := &Handler{}
	switch cfg.Provider {
	case "gitlab":
		h.Handler = &gitlab.GitlabHandler{
			Exporters: exporters,
		}
	}
	return h, nil
}

func (h *Handler) MergeRequestHandlerFunc(ctx *fasthttp.RequestCtx) {
	_, err := h.Handler.HandleMergeRequest(ctx.Request.Body())
	if err != nil {
		h.processError(ctx, err)
		return
	}
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (h *Handler) DeploymentHandlerFunc(ctx *fasthttp.RequestCtx) {
	_, err := h.Handler.HandleDeployment(ctx.Request.Body())
	if err != nil {
		h.processError(ctx, err)
		return
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
