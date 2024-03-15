package uvalidator

import (
	"github.com/stdyum/api-common/errors"
)

var (
	ErrSlice = errors.WrapString(errors.ErrValidation, "slice")
)

func SliceNotEmpty[T any](prop string, value []T) error {
	if len(value) == 0 {
		return errors.WrapStringf(ErrSlice, "%s must not be an empty array", prop)
	}
	return nil
}
