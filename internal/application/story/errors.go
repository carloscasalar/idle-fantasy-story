package story

import (
	"github.com/carloscasalar/idle-fantasy-story/internal/application"
)

var (
	// ErrWorldIDIsRequired is returned when a world ID is not provided
	ErrWorldIDIsRequired = application.NewValidationError("world ID is required to create a new story")
)
