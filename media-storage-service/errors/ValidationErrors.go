package errors

import (
	"errors"
	"fmt"
)

var (
	ErrUUIDInvalidLenght = errors.New("uuid validation failed")
	ErrInvalidUUID       = fmt.Errorf("uuid validation failed")
)
