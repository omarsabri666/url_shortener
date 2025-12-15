package errs

import "fmt"

type AppError struct {
	Code        int    `json:"-"`
	Message     string `json:"message"`
	Details     string `json:"details"`
	InternalErr error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}
func New(code int, msg string, details string, internal error) *AppError {
	return &AppError{
		Code:        code,
		Message:     msg,
		Details:     details,
		InternalErr: internal,
	}
}
