package story

import "errors"

var (
	// ErrWorldIDIsRequired is returned when a world ID is not provided
	ErrWorldIDIsRequired = errors.New("world ID is required to create a new story")
)
