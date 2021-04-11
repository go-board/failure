package failure

import (
	"fmt"
)

type Errno int32

type Error struct {
	inner error
	frame *Frame

	Msg  string
	Code Errno
}

func (e *Error) Error() string {
	if e.inner != nil {
		if e.frame != nil {
			return fmt.Sprintf("location: [%s] code: %d, msg: %s, err: %s", e.frame.Location, e.Code, e.Msg, e.inner)
		}
		return fmt.Sprintf("code: %d, msg: %s, err: %s", e.Code, e.Msg, e.inner)
	} else {
		if e.frame != nil {
			return fmt.Sprintf("location: [%s] code: %d, msg: %s", e.frame.Location, e.Code, e.Msg)
		}
		return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
	}
}

func (e *Error) Unwrap() error { return e.inner }

func New(code Errno, msg string) *Error {
	frames := traceback(0, 1)
	var frame *Frame
	if len(frames) > 0 {
		frame = frames[0]
	}
	return &Error{Code: code, Msg: msg, frame: frame}
}

func (e *Error) WithError(err error) *Error {
	e.inner = err
	return e
}

func (e *Error) WithFrame(f *Frame) *Error {
	e.frame = f
	return e
}

func (e *Error) Frame() *Frame {
	var f *Frame
	*f = *e.frame
	return f
}
