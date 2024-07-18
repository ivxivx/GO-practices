package domain

import "fmt"

type NotFoundError struct {
	ID string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("resource with ID %s not found", e.ID)
}
