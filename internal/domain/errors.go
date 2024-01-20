package domain

type ErrorType int

var (
	// ErrWorldDoesNotExist is returned when a world does not exist
	ErrWorldDoesNotExist = NewResourceDoesNotExistError("world does not exist")
)

const (
	UnexpectedError ErrorType = iota
	ValidationError
	LogicError
	ResourceDoesNotExistError
)

type Error struct {
	errorType ErrorType
	message   string
}

func (e *Error) Error() string {
	return e.message
}

func NewUnexpectedError(message string) *Error {
	return &Error{
		errorType: UnexpectedError,
		message:   message,
	}
}

func NewValidationError(message string) *Error {
	return &Error{
		errorType: ValidationError,
		message:   message,
	}
}

func NewLogicError(message string) *Error {
	return &Error{
		errorType: LogicError,
		message:   message,
	}
}

func NewResourceDoesNotExistError(message string) *Error {
	return &Error{
		errorType: ResourceDoesNotExistError,
		message:   message,
	}
}
