package application

import "errors"

var (
	// ErrInternalServer is returned when an unexpected error occurs
	ErrInternalServer = errors.New("internal server error")
)
