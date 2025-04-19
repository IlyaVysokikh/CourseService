package errors

import (
	"errors"
	"fmt"
)

// Общие типы ошибок
var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrInternal     = errors.New("internal error")
)

// Типизированная ошибка, которая несет контекст
type Error struct {
	Code    error
	Message string 
	Err     error  
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s", e.Message)
	}
	return fmt.Sprintf("%s",  e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func New(code error, msg string, err error) *Error {
	return &Error{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, ErrInvalidInput)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}
