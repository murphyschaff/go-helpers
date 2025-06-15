package helpers

import (
	"fmt"
)

const FilepathError = "Invalid filepath entry. The file was entered wrong, or does not exist on the system."

type StringError struct {
	arg     string
	message string
}

func (e *StringError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.message)
}

type IntError struct {
	arg     int
	message string
}

func (e *IntError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}
