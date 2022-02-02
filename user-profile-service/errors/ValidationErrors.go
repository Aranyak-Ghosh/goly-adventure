package errors

import (
	"errors"
	"fmt"
)

var (
	ErrUUIDLenght  = errors.New("uuid validation failed")
	ErrInvalidUUID = fmt.Errorf("uuid validation failed %w", ErrUUIDLenght)
)
