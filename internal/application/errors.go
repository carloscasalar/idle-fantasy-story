package application

var (
	// ErrInternalServer is returned when an unexpected error occurs
	ErrInternalServer = NewUnexpectedError("internal server error")
)

type ErrorType int

const (
	UnexpectedError ErrorType = iota
	InvalidRequest
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

func NewInvalidRequestError(message string) *Error {
	return &Error{
		errorType: InvalidRequest,
		message:   message,
	}
}

func NewLogicError(message string) *Error {
	return &Error{
		errorType: LogicError,
		message:   message,
	}
}
