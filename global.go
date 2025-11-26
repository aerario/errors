package errors

import (
	"errors"
	"hash/fnv"
)

// TODO: adapters (sql, validation etc)

func KindOf(err error) Kind {
	if t, ok := err.(*implementation); ok {
		return t.kind
	}

	return ErrKindGeneral
}

func From(err error) Error {
	if err == nil {
		return nil
	}

	if t, ok := err.(*implementation); ok {
		return t
	}

	return &implementation{
		kind:     ErrKindGeneral,
		message:  err.Error(),
		previous: errors.Unwrap(err),
	}
}

func Is(err error, target any) bool {
	switch t := target.(type) {
	case Factory:
		// since errors will compare factory and factory's error by implementation.Id,
		// we can just call new with no arguments
		return errors.Is(err, t.New())

	case error:
		return errors.Is(err, t)
	}

	return false
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func Labels(err error) LabelList {
	if t, ok := err.(*implementation); ok {
		return t.labels
	}

	return LabelList{}
}

func IsUserFriendly(err error) bool {
	return Labels(err).Has(LabelUserFriendly)
}

func In(err error, target ...any) bool {
	for _, t := range target {
		if Is(err, t) {
			return true
		}
	}

	return false
}

// As is a wrapper over std function to avoid multiple errors package imports
func As(err error, target any) bool {
	return errors.As(err, target)
}

func errorId(template string) uint32 {
	var hash = fnv.New32()
	// this function never returns error, so we can ignore it
	_, _ = hash.Write([]byte(template))
	return hash.Sum32()
}
