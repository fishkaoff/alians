package errs

import "errors"

var (
	InternalError = errors.New("internal error")

	EmptyName  = errors.New("name cannot be empty")
	EmptyPhone = errors.New("phone cannot be empty")
	EmptyEmail = errors.New("email cannot be empty")
)
