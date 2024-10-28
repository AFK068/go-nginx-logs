package domain

import "fmt"

type MissedMandatoryFlagError struct {
	Message string
}

func (e *MissedMandatoryFlagError) Error() string {
	return fmt.Sprintf("Missed mandatory flag: %s", e.Message)
}
