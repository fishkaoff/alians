package validate

import (
	"testing"

	"github.com/fishkaoff/alians/notificator/notificator/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateMessage(t *testing.T) {
	validMsg := &models.Message{
		Name:  "John Doe",
		Phone: "+123456789",
		Email: "john@example.com",
	}

	emptyNameMsg := &models.Message{
		Name:  "",
		Phone: "+123456789",
		Email: "john@example.com",
	}

	emptyPhoneMsg := &models.Message{
		Name:  "John Doe",
		Phone: "",
		Email: "john@example.com",
	}

	emptyEmailMsg := &models.Message{
		Name:  "John Doe",
		Phone: "+123456789",
		Email: "",
	}

	// Test valid message
	err := ValidateMessage(validMsg)
	assert.Nil(t, err, "Expected no error for a valid message")

	// Test empty name
	err = ValidateMessage(emptyNameMsg)
	assert.EqualError(t, err, "name cannot be empty", "Expected error for empty name")

	// Test empty phone
	err = ValidateMessage(emptyPhoneMsg)
	assert.EqualError(t, err, "phone cannot be empty", "Expected error for empty phone")

	// Test empty email
	err = ValidateMessage(emptyEmailMsg)
	assert.EqualError(t, err, "email cannot be empty", "Expected error for empty email")
}
