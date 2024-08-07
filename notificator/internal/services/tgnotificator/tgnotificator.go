package tgnotificator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fishkaoff/alians/notificator/notificator/internal/config"
	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
	"github.com/fishkaoff/alians/notificator/notificator/internal/lib/errs"
)

var tgurl = "https://api.telegram.org/bot%s/sendMessage"

type TgNotificator struct {
	BotToken string
	ChatID   string

	log *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *TgNotificator {
	return &TgNotificator{
		BotToken: cfg.TgNotificatorConfig.BotToken,
		ChatID:   cfg.TgNotificatorConfig.ChatID,
		log:      log,
	}
}

func (n *TgNotificator) ThrowMessage(ctx context.Context, msg *models.Message) error {

	body, err := n.prepareRequestBody(msg)
	if err != nil {
		n.log.Error("cannot prepare request body", slog.String("error", err.Error()))
		return errs.ErrInternalError
	}

	url := fmt.Sprintf(tgurl, n.BotToken)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		n.log.Error("error while request to telegram: ", slog.String("error", err.Error()), slog.String("op", "notificator.ThrowMessage"))
		return errs.ErrInternalError
	}
	defer resp.Body.Close()

	return nil
}

func (n *TgNotificator) renderMessage(msg *models.Message) string {
	return fmt.Sprintf("*-Новая заявка:*\n\n*-Система: *%s\n*-Телефон: *%s\n*-Площадь: *%s", msg.System, msg.Phone, msg.Square)
}

func (n *TgNotificator) prepareRequestBody(msg *models.Message) ([]byte, error) {
	text := n.renderMessage(msg)
	reqBody := models.RequestBody{
		ChatID:    n.ChatID,
		ParseMode: "Markdown",
		Text:      text,
	}

	encodedReqBody, err := json.Marshal(reqBody)
	if err != nil {
		n.log.Error("error to encode body", slog.String("error", err.Error()))
		return []byte(""), err
	}

	return encodedReqBody, nil
}
