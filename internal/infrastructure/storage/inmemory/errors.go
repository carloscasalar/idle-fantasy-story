package inmemory

import "errors"

var (
	// ErrUnableToParseYmlWorlds is returned when the yml file containing the worlds cannot be parsed
	ErrUnableToParseYmlWorlds = errors.New("unable to parse yml worlds")
	// ErrUnableToRetrieveWorlds is returned when the worlds cannot be retrieved from the store
	ErrUnableToRetrieveWorlds = errors.New("unable to retrieve worlds from the store")
	// ErrInvalidSpecies is returned when an invalid species is found
	ErrInvalidSpecies = errors.New("unable to parse worlds: invalid species")
)
