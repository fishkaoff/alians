package validate

import (
	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
	"github.com/fishkaoff/alians/notificator/notificator/internal/lib/errs"
)

func ValidateMessage(msg *models.Message) error {
	if msg.From == "" {
		return errs.ErrEmptyFrom
	}

	if msg.Name == "" {
		return errs.ErrEmptyName
	}

	if msg.Phone == "" {
		return errs.ErrEmptyPhone
	}

	if msg.Email == "" {
		return errs.ErrEmptyEmail
	}

	return nil
}
