package utils

import "fmt"

func AppendError(existErr, NewErr error) error {
	if existErr == nil {
		return NewErr
	}
	return fmt.Errorf("%v, %w", existErr, NewErr)
}
