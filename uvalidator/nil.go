package uvalidator

import (
	"github.com/stdyum/api-common/errors"
)

var (
	ErrNil = errors.WrapString(errors.ErrValidation, "nil")
)

func IsNotNil(prop string, value any) error {
	if value != nil {
		return nil
	}

	return errors.WrapStringf(ErrNil, "%s must not be nil", prop)
}
