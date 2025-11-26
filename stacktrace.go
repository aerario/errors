package lapsus

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"fmt"
	"runtime"
)

type Stacker interface {
	StackTrace() jsontext.Value
	Location() string

	error
}

type stackTraceFrame struct {
	Kind     string    `json:"kind"`
	Labels   LabelList `json:"labels"`
	Error    string    `json:"error"`
	Location string    `json:"location,omitempty"`
}

type location struct {
	file string
	line int
}

func (loc location) String() string {
	return fmt.Sprintf("%s:%d", loc.file, loc.line)
}

func (self *implementation) setLocation(callDepth int) {
	_, self.location.file, self.location.line, _ = runtime.Caller(callDepth + 1)
}

func (self *implementation) Location() string {
	return self.location.String()
}

func (self *implementation) StackTrace() jsontext.Value {
	var (
		out = make([]stackTraceFrame, 0, 1)
		err = error(self)
	)

	for err != nil {
		var frame = stackTraceFrame{
			Kind:   kindName(KindOf(err)),
			Labels: Labels(err),
		}

		if t, ok := err.(fmt.Stringer); ok {
			frame.Error = t.String()
		} else {
			frame.Error = err.Error()
		}

		if t, ok := err.(Stacker); ok {
			frame.Location = t.Location()
		}

		out = append(out, frame)
		err = errors.Unwrap(err)
	}

	var js []byte
	if js, err = json.Marshal(out); err != nil {
		return jsontext.Value(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
	}

	return js
}
