package httpserver

import (
	"context"
	"log/slog"

	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
	"github.com/fishkaoff/alians/notificator/notificator/internal/lib/validate"
	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	app *fiber.App

	ListenAddr string
	log        *slog.Logger

	msgThrower MessageThrower
}

func New(ListenAddr string, log *slog.Logger, msgThrower MessageThrower) *HttpServer {
	return &HttpServer{
		ListenAddr: ListenAddr,
		log:        log,
		msgThrower: msgThrower,
	}
}

type MessageThrower interface {
	ThrowMessage(ctx context.Context, message models.Message)
}

func (hs *HttpServer) MustStart() {
	hs.app = fiber.New()
	hs.setupEndpoints()

	hs.log.Info("server is running on: ", slog.String("addr", hs.ListenAddr))
	err := hs.app.Listen(hs.ListenAddr)
	if err != nil {
		panic(err)
	}
}

func (hs *HttpServer) setupEndpoints() {
	message := hs.app.Group("/message")

	message.Get("/new", hs.newMessageHandler)
}

func (hs *HttpServer) newMessageHandler(c *fiber.Ctx) error {
	var msg models.Message
	if err := c.BodyParser(&msg); err != nil {
		hs.log.Error(err.Error())
		return c.Status(400).JSON(fiber.Map{
			"response": "bad request",
		})
	}

	hs.log.Debug("handled new message: ", &msg)

	if err := validate.ValidateMessage(msg); err != nil {
		hs.log.Debug("message not valid: ", err)
		return c.Status(400).JSON(fiber.Map{
			"response": err,
		})
	}

	hs.log.Debug("valid message")

	go hs.msgThrower.ThrowMessage(context.TODO(), msg)

	return c.Status(200).JSON(fiber.Map{
		"response": "success",
	})
}