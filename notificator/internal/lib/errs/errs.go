package errs

import "errors"

var (
	ErrInternalError = errors.New("internal error")

	ErrEmptyName  = errors.New("name cannot be empty")
	ErrEmptyPhone = errors.New("phone cannot be empty")
	ErrEmptyEmail = errors.New("email cannot be empty")
	ErrEmptyFrom = errors.New("from field cannot be empty")
)
