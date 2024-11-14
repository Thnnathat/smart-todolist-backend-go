package errors

import "errors"

var (
	ErrNotfound       = errors.New("not found")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternal       = errors.New("internal server error")
)
