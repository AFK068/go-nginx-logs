package pathutils

import "fmt"

type PathError struct {
	Message string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("Path error: %s", e.Message)
}
