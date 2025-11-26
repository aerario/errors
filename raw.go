package lapsus

import (
	"encoding/json/jsontext"
	"errors"
	"strings"
)

func Raw(err error) error {
	if t, ok := err.(*implementation); ok {
		return &raw{
			err: t,
		}
	}

	return err
}

type raw struct {
	err *implementation
}

func (self *raw) Error() string {
	var (
		err = error(self.err)
		out []string
	)

	for err != nil {
		if t, ok := err.(*implementation); ok {
			out = append(out, t.String())
		} else {
			out = append(out, err.Error())
		}

		err = errors.Unwrap(err)
	}

	if len(out) == 0 {
		return DefaultUserFriendlyError
	}

	return strings.Join(out, ": ")
}

func (self *raw) Location() string {
	return self.err.Location()
}

func (self *raw) StackTrace() jsontext.Value {
	return self.err.StackTrace()
}

func (self *raw) Unwrap() error {
	return self.err
}
