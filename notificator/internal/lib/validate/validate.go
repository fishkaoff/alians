package validate

import (
	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
	"github.com/fishkaoff/alians/notificator/notificator/internal/lib/errs"
)

func ValidateMessage(msg *models.Message) error {
	if msg.System == "" {
		return errs.ErrEmptySystem
	}

	if msg.Square == "" {
		return errs.ErrEmptySquare
	}

	if msg.Phone == "" {
		return errs.ErrEmptyPhone
	}

	return nil
}
