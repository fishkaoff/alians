package errs

import "errors"

var (
	ErrInternalError = errors.New("internal error")

	ErrEmptySystem  = errors.New("system cannot be empty")
	ErrEmptyPhone = errors.New("phone cannot be empty")
	ErrEmptySquare = errors.New("square cannot be empty")
)
