package domain

import "fmt"

type MissedMandatoryFlagError struct {
	Message string
}

func (e *MissedMandatoryFlagError) Error() string {
	return fmt.Sprintf("Missed mandatory flag: %s", e.Message)
}

type PathError struct {
	Message string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("Path error: %s", e.Message)
}

type ParseNGINXStringError struct {
	Message string
}

func (e *ParseNGINXStringError) Error() string {
	return fmt.Sprintf("Parse NGINX string error: %s", e.Message)
}

type InvalidDateRangeError struct {
	Message string
}

func (e *InvalidDateRangeError) Error() string {
	return fmt.Sprintf("Invalid date range: %s", e.Message)
}
