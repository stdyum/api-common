package uvalidator

import (
	"github.com/stdyum/api-common/errors"
	"github.com/stdyum/api-common/uslices"
)

func V(results ...error) error {
	errs := uslices.FilterFunc(results, func(item error) bool {
		return item != nil
	})

	if len(errs) == 0 {
		return nil
	}

	return errors.WrapJoin(errors.ErrValidation, errs...)
}
