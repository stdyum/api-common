package uvalidator

import (
	"github.com/stdyum/api-common/errors"
)

var (
	ErrString = errors.WrapString(errors.ErrValidation, "string")
)

func StringIn(prop string, value string, in ...string) error {
	for _, v := range in {
		if v == value {
			return nil
		}
	}
	return errors.WrapStringf(ErrString, "%s must be one of %v", prop, in)
}
