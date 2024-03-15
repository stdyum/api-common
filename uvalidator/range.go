package uvalidator

import (
	"github.com/stdyum/api-common/errors"
)

var (
	ErrRange = errors.WrapString(errors.ErrValidation, "range")
)

func RangeInt(prop string, value int, from, to int) error {
	if value < from || value > to {
		return errors.WrapStringf(ErrRange, "%s must be in bounds (%d...%d)", prop, from, to)
	}
	return nil
}
