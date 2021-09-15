package failure

import (
	"fmt"
)

type Errno int32

type Error struct {
	inner error
	frame *Frame
	Msg   string
	Code  Errno
}

func New(code Errno, msg string) *Error {
	frames := traceback(0, 2)
	var frame *Frame
	if len(frames) > 1 {
		frame = frames[1]
	}
	return &Error{Code: code, Msg: msg, frame: frame}
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

func (e *Error) WithError(err error) *Error {
	return &Error{inner: err, frame: e.frame, Msg: e.Msg, Code: e.Code}
}

func (e *Error) WithFrame(f *Frame) *Error {
	return &Error{inner: e.inner, frame: f, Msg: e.Msg, Code: e.Code}
}

func (e *Error) Frame() *Frame {
	if e.frame == nil {
		return nil
	}
	return &Frame{Location: e.frame.Location, FuncName: e.frame.FuncName}
}
