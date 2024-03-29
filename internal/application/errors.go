package application

var (
	// ErrInternalServer is returned when an unexpected error occurs
	ErrInternalServer    = NewUnexpectedError("internal server error")
	ErrWorldDoesNotExist = NewValidationError("world does not exist")
)

type ErrorType int

const (
	UnexpectedError ErrorType = iota
	ValidationError
	LogicError
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
