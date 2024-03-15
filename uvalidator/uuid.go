package uvalidator

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/errors"
)

var (
	ErrUUID = errors.WrapString(errors.ErrValidation, "uuid")
)

func UUIDNotNil(prop string, value uuid.UUID) error {
	if value == uuid.Nil {
		return errors.WrapStringf(ErrUUID, "%s is nil", prop)
	}
	return nil
}
