package story

import (
	"github.com/carloscasalar/idle-fantasy-story/internal/application"
)

var (
	// ErrWorldIDIsRequired is returned when a world ID is not provided
	ErrWorldIDIsRequired = application.NewValidationError("world ID is required to create a new story")
	// ErrInvalidPartySize is returned when the party size is set and is not between 3 and 6
	ErrInvalidPartySize = application.NewValidationError("party size, if any, must be between 3 and 6 characters")
)
