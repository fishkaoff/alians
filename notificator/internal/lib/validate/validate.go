package validate

import (
	"errors"

	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
)

func ValidateMessage(msg models.Message) error {
	if msg.Name == "" {
		return errors.New("name cannot be empty")
	}

	if msg.Phone == "" {
		return errors.New("phone cannot be empty")

	}

	if msg.Email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}