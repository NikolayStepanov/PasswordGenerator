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

type App struct {
	config     *config.Config
	mux        muxer
	server     *server.Server
	services   *service.Services
	repository *repository.Repository
}

func RegisterPasswordHandlers(config *config.Config, mux *http.ServeMux, handlerResponder *handler.Handler) {
}

func NewApp(config *config.Config) (*App, error) {
	mux := http.NewServeMux()
	mux.HandleFunc(config.PathHandles.Index, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	server := server.NewServer(config, mux)

	return &App{
		config: config,
		mux:    mux,
		server: server,
	}, nil
}

func Run() {

	cfg := config.Init()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	//HTTP Server
	app, err := NewApp(cfg)
	if err != nil {
		log.Fatalf("error new app: %s \n", err)
	}
	go func() {
		defer cancel()
		if err := app.server.Run(); err != nil {
			log.Printf("error occurred while running http server: %s \n", err.Error())
		}
	}()

	log.Println("password generator is running")
	log.Println("server started")

	<-ctx.Done()

	log.Printf("shutting down server\n")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelShutdown()

	if err = app.server.Stop(ctxShutdown); err != nil {
		log.Printf("error occured on server shutting down: %s \n", err.Error())
	}
	log.Println("password generator stopped")
}
