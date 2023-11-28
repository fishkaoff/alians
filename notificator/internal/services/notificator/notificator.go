package notificator

import (
	"context"
	"log/slog"

	"github.com/fishkaoff/alians/notificator/notificator/internal/config"
	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
)

type Notificator struct {
	BotToken string
	ChatID   string

	log *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *Notificator {
	return &Notificator{
		BotToken: cfg.NotificatorConfig.BotToken,
		ChatID:   cfg.NotificatorConfig.ChatID,
		log: log,
	}
}

func (n *Notificator) ThrowMessage(ctx context.Context, msg models.Message) {
	panic("implement me")
}
