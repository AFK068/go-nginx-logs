package application

import "fmt"

type InvalidRenderFormatError struct {
	Message string
}

func (e *InvalidRenderFormatError) Error() string {
	return fmt.Sprintf("Invalid render format: %s", e.Message)
}

type InvalidPathFormatError struct {
	Message string
}

func (e *InvalidPathFormatError) Error() string {
	return fmt.Sprintf("Invalid path: %s", e.Message)
}
