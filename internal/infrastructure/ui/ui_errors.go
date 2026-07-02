package ui

import "fmt"

type InputError struct {
	msg string
	err error
}

func (inputErr InputError) Error() string {
	return fmt.Sprintf("%s: %v", inputErr.msg, inputErr.err)
}

func (inputErr InputError) Unwrap() error {
	return inputErr.err
}

type OutputError struct {
	msg string
	err error
}

func (outputErr OutputError) Error() string {
	return fmt.Sprintf("%s: %v", outputErr.msg, outputErr.err)
}

func (outputErr OutputError) Unwrap() error {
	return outputErr.err
}
