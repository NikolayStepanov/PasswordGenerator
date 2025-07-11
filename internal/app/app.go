// Package app initializes and runs the main application components,
// including configuration, server, routing, services, and data access
package app

import (
	"context"
	"fmt"
	"github.com/NikolayStepanov/PasswordGenerator/internal/delivery/http/handler"
	"github.com/NikolayStepanov/PasswordGenerator/internal/repository"
	"github.com/NikolayStepanov/PasswordGenerator/internal/server"
	"github.com/NikolayStepanov/PasswordGenerator/internal/service"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikolayStepanov/PasswordGenerator/internal/config"
)

type muxer interface {
	Handle(pattern string, handler http.Handler)
}

// App represents the main application structure
type App struct {
	cnf        *config.Config
	mux        muxer
	serverHTTP *server.Server
	services   *service.Services
	repository *repository.Repository
}

// RegisterPasswordHandlers registers HTTP handlers related to password operations
func RegisterPasswordHandlers(cnf *config.Config, mux *http.ServeMux, handlerResp *handler.Handler) {
	mux.Handle(cnf.PathHandles.Password, handler.NewGetPasswordHandler(cnf.PathHandles.Password, handlerResp))
}

// NewApp creates and initializes a new application based on the provided configuration
func NewApp(cnf *config.Config) (*App, error) {
	mux := http.NewServeMux()
	mux.HandleFunc(cnf.PathHandles.Password, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	serverHTTP := server.NewServer(cnf, mux)

	return &App{
		cnf:        cnf,
		mux:        mux,
		serverHTTP: serverHTTP,
	}, nil
}

// Run launching the application
func Run() {
	cfg := config.Init()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	// HTTP Server
	app, err := NewApp(cfg)
	if err != nil {
		log.Panicf("error new app: %s \n", err)
	}

	go func() {
		defer cancel()
		if errApp := app.serverHTTP.Run(); errApp != nil {
			log.Printf("error occurred while running http server: %s \n", errApp.Error())
		}
	}()

	log.Println("password generator is running")
	log.Println("server started")

	<-ctx.Done()

	log.Printf("shutting down server\n")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelShutdown()

	if err = app.serverHTTP.Stop(ctxShutdown); err != nil {
		log.Printf("error occurred on server shutting down: %s \n", err.Error())
	}
	log.Println("password generator stopped")
}
