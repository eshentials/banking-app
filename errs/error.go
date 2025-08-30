package errs

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: 404, Message: message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: 500, Message: message}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}
