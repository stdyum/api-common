package uvalidator

import (
	"time"

	"github.com/stdyum/api-common/errors"
)

var (
	ErrTime = errors.WrapString(errors.ErrValidation, "time")
)

func TimeBefore(prop string, value time.Time, t time.Time) error {
	if !value.Before(t) {
		return errors.WrapStringf(ErrTime, "%s must be before %s", prop, t)
	}
	return nil
}

func TimeBeforeOrEqual(prop string, value time.Time, t time.Time) error {
	if value.After(t) {
		return errors.WrapStringf(ErrTime, "%s must be before or equal %s", prop, t)
	}
	return nil
}

func TimeAfter(prop string, value time.Time, t time.Time) error {
	if !value.After(t) {
		return errors.WrapStringf(ErrTime, "%s must be after %s", prop, t)
	}
	return nil
}

func TimeAfterOrEqual(prop string, value time.Time, t time.Time) error {
	if value.Before(t) {
		return errors.WrapStringf(ErrTime, "%s must be after or equal %s", prop, t)
	}
	return nil
}

func DurationBefore(prop string, value time.Duration, d time.Duration) error {
	if value >= d {
		return errors.WrapStringf(ErrTime, "%s must be before %s", prop, d)
	}
	return nil
}

func DurationBeforeOrEqual(prop string, value time.Duration, d time.Duration) error {
	if value > d {
		return errors.WrapStringf(ErrTime, "%s must be before or equal %s", prop, d)
	}
	return nil
}

func DurationAfter(prop string, value time.Duration, d time.Duration) error {
	if value <= d {
		return errors.WrapStringf(ErrTime, "%s must be after %s", prop, d)
	}
	return nil
}

func DurationAfterOrEqual(prop string, value time.Duration, d time.Duration) error {
	if value < d {
		return errors.WrapStringf(ErrTime, "%s must be after or equal %s", prop, d)
	}
	return nil
}

func DateBefore(prop string, value time.Time, t time.Time) error {
	return TimeBefore(prop, truncateTimePart(value), truncateTimePart(t))
}

func DateBeforeOrEqual(prop string, value time.Time, t time.Time) error {
	return TimeBeforeOrEqual(prop, truncateTimePart(value), truncateTimePart(t))
}

func DateAfter(prop string, value time.Time, t time.Time) error {
	return TimeAfter(prop, truncateTimePart(value), truncateTimePart(t))

}

func DateAfterOrEqual(prop string, value time.Time, t time.Time) error {
	return TimeAfterOrEqual(prop, truncateTimePart(value), truncateTimePart(t))
}

func truncateTimePart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
