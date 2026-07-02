package session

import "fmt"

type AppError struct {
	msg string
	err error
}

func (appErr AppError) Error() string {
	return fmt.Sprintf("%s: %v", appErr.msg, appErr.err)
}

func (appErr AppError) Unwrap() error {
	return appErr.err
}
