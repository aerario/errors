package errors

import (
	"fmt"
)

type Factory interface {
	WithLabels(...Label) Factory
	New(args ...interface{}) Error
}

func NewFactory(kind Kind, template string) Factory {
	return &factory{
		id:       errorId(template),
		kind:     kind,
		template: template,
		labels:   LabelList{LabelUserFriendly}, // factory-made errors are always user-friendly
	}
}

type factory struct {
	// id of a factory is a hash of its template
	// we cannot just compare error messages to find out whether the error is of type X
	// because of the arguments, so we need to remember its template.
	id       uint32
	kind     Kind
	template string
	labels   LabelList
}

func (f factory) New(args ...interface{}) Error {
	var err = &implementation{
		id:      f.id,
		kind:    f.kind,
		message: fmt.Sprintf(f.template, args...),
	}

	err.setLocation(1)

	return err.WithLabels(f.labels...)
}

func (f factory) WithLabels(labels ...Label) Factory {
	var newFactory = f

	newFactory.labels = make(LabelList, 0, len(f.labels)+len(labels))
	newFactory.labels = append(newFactory.labels, f.labels...)
	newFactory.labels = append(newFactory.labels, labels...)

	return &newFactory
}
