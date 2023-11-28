package app

import (
	"log/slog"

	"github.com/fishkaoff/alians/notificator/notificator/internal/config"
	"github.com/fishkaoff/alians/notificator/notificator/internal/httpserver"
	"github.com/fishkaoff/alians/notificator/notificator/internal/services/notificator"
)

type App struct {
	HttpServer *httpserver.HttpServer
}

func New(cfg *config.Config, log *slog.Logger) *App {
	ntfSvc := notificator.New(cfg, log)

	server := httpserver.New(cfg.HttpConfig.ListenAddr, log, ntfSvc)
	return &App{
		HttpServer: server,
	} 
}
