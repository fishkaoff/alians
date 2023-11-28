package validate

import (
	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
	"github.com/fishkaoff/alians/notificator/notificator/internal/lib/errs"
)

func ValidateMessage(msg *models.Message) error {
	if msg.Name == "" {
		return errs.EmptyName
	}

	if msg.Phone == "" {
		return errs.EmptyPhone
	}

	if msg.Email == "" {
		return errs.EmptyEmail
	}

	return nil
}
