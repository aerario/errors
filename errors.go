package errors

import (
	"errors"
	"fmt"
	"maps"
	"strings"
)

var DefaultUserFriendlyError = "something went wrong"

type Error interface {
	Annotate(message string, args ...interface{}) Error
	Wrap(err error) Error
	Unwrap() error
	Labels() LabelList
	WithLabels(...Label) Error
	Details() map[string]string

	error
}

func New(kind Kind, message string, args ...interface{}) Error {
	var err = &implementation{
		id:      errorId(message),
		kind:    kind,
		message: fmt.Sprintf(message, args...),
	}

	err.setLocation(2)

	return err
}

type implementation struct {
	// id of an error is a hash of its template
	// we cannot just compare error messages to find out whether the error if of type X
	// because of the arguments, so we need to remember its template.
	id       uint32
	kind     Kind
	labels   LabelList
	message  string
	previous error
	details  map[string]string
	location location
}

func (self *implementation) Annotate(message string, args ...interface{}) Error {
	self.message += ": " + fmt.Sprintf(message, args...)
	return self
}

func (self *implementation) Wrap(err error) Error {
	self.previous = err
	return self
}

func (self *implementation) Unwrap() error {
	return self.previous
}

func (self *implementation) Labels() LabelList {
	return self.labels
}

func (self *implementation) WithLabels(in ...Label) Error {
	self.labels = self.labels.Add(in...)
	return self
}

func (self *implementation) Details() map[string]string {
	var details = make(map[string]string)

	var prev Error
	if As(self.previous, &prev) {
		details = prev.Details()
	}

	maps.Copy(details, self.details)

	return details
}

// External interface implementations

func (self *implementation) Is(target error) bool {
	switch target := target.(type) {
	case *implementation:
		if self.id == target.id && self.kind == target.kind {
			return true
		}
	default:
		if self.message == target.Error() {
			return true
		}
	}

	return false
}

func (self *implementation) String() string {
	return self.message
}

func (self *implementation) ErrorCode() string {
	return kindName(self.kind)
}

// Error concatenates and prints out all underlying user-friendly errors
func (self *implementation) Error() string {
	var (
		err = error(self)
		out []string
	)

	for err != nil {
		switch t := err.(type) {
		case *implementation:
			if t.labels.Has(LabelUserFriendly) {
				out = append(out, t.String())
			}
		}

		err = errors.Unwrap(err)
	}

	if len(out) == 0 {
		return DefaultUserFriendlyError
	}

	return strings.Join(out, ": ")
}

func (self implementation) MarshalJSON() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%q", self.Error()), nil
}
