package handler

import (
	"fmt"
	"net/http"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	"github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Handler port.VCSHostingHandler
}

func New(cfg *config.Handler) (*Handler, error) {
	handler := &Handler{}
	h, err := vcshosting.GetVCSHostingHandler(cfg)
	if err != nil {
		return nil, err
	}
	handler.Handler = h
	return handler, nil
}

func (h *Handler) MergeRequestHandlerFunc(ctx *fasthttp.RequestCtx) {
	if err := h.Handler.HandleMergeRequest(ctx.Request.Body()); err != nil {
		h.processError(ctx, err)
		return
	}
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (h *Handler) DeploymentHandlerFunc(ctx *fasthttp.RequestCtx) {
	if err := h.Handler.HandleDeployment(ctx.Request.Body()); err != nil {
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
