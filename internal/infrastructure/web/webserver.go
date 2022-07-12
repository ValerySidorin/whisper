package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ValerySidorin/whisper/internal/infrastructure/appctx"
	"github.com/ValerySidorin/whisper/internal/infrastructure/config"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Server struct {
	Server    fasthttp.Server
	Stop      chan bool
	IsStarted atomic.Value
	Router    *router.Router
	Config    config.HTTP
	Ctx       context.Context
	Cc        *appctx.CoreContext
}

func Register(
	ctx *appctx.CoreContext,
	routes *router.Router,
	cfg *config.Configuration,
) *Server {
	webServer := new(Server)
	webServer.Router = routes
	webServer.IsStarted.Store(false)
	webServer.Stop = make(chan bool)
	webServer.Config = cfg.Http
	webServer.Ctx = ctx.Ctx()
	webServer.Cc = ctx
	return webServer
}

func (w *Server) ServerStop() {
	if w.IsStarted.Load().(bool) {
		w.Stop <- true
	}
}

func (w *Server) IsServerStarted() bool {
	return w.IsStarted.Load().(bool)
}

func (w *Server) Serve() {
	w.Server = fasthttp.Server{
		Handler:            w.Router.Handler,
		ReadTimeout:        w.Config.GetTimeout(),
		WriteTimeout:       w.Config.GetTimeout(),
		DisableKeepalive:   true,
		TCPKeepalive:       false,
		MaxRequestsPerConn: 1,
		Name:               config.ProjectName,
	}
	go func() {
		w.IsStarted.Store(true)
		err := w.Server.ListenAndServe(":" + w.Config.Port)
		if err != nil {
			if err != http.ErrServerClosed {
				fmt.Println("Can't start webserver: " + err.Error())
			} else {
				fmt.Println(err)
			}
		}
	}()

	go func() {
		<-w.Stop
		log.Println("Webserver stop signal received")
		shutDownCtx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		done := make(chan struct{})
		go func() {
			err := w.Server.Shutdown()
			if err != nil {
				log.Fatal("Webserver shutdown error: " + err.Error())
			}
			done <- struct{}{}
		}()
		select {
		case <-shutDownCtx.Done():
			log.Println("Webserver shutdown forced")
		case <-done:
			log.Println("Webserver shutdown completed")
		}
		cancel()
		close(w.Stop)
		w.IsStarted.Store(false)
	}()
}

func (w *Server) Run() {
	w.Serve()

	//graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	stopChannel := make(chan bool)

	go func(ch <-chan os.Signal, st chan<- bool) {
		<-ch
		log.Println("Stop signal received")
		w.ServerStop()
		log.Println("Stop signal sent to webserver")
		w.Cc.Cancel()
		log.Println("Waiting while webserver is stopping")
		for w.IsServerStarted() {
			time.Sleep(time.Microsecond)
		}
		log.Println("Webserver stopped")
		st <- true
	}(signalChannel, stopChannel)

	<-stopChannel
}
