package errors

import (
	"errors"
	"fmt"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrValidation   = errors.New("validation")
)

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func Wrap(target error, err error) error {
	return WrapString(target, err.Error())
}

func WrapString(target error, err string) error {
	return fmt.Errorf("%s: %w", err, target)
}

func WrapStringf(target error, err string, args ...any) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(err, args...), target)
}

func WrapJoin(target error, errs ...error) error {
	return WrapString(target, errors.Join(errs...).Error())
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}
