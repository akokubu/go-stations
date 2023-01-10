package model

import "fmt"

type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("not found error")
}
