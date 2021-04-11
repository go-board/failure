package failure

import (
	"fmt"
	"runtime"
)

const maxCallerDepth = 256

type Frame struct {
	Location string
	FuncName string
}

func (f *Frame) GetLoction() string {
	if f != nil {
		return f.Location
	}
	return "???"
}

func (f *Frame) GetFuncName() string {
	if f != nil {
		return f.FuncName
	}
	return "???"
}

type Frames []*Frame

func traceback(skip int, depth int) Frames {
	var frames Frames
	if depth <= 0 {
		depth = maxCallerDepth
	}
	pc := make([]uintptr, depth)
	n := runtime.Callers(skip+2, pc)
	for i := 0; i < n; i++ {
		fn := runtime.FuncForPC(pc[i])
		f, l := fn.FileLine(pc[i])
		frames = append(frames, &Frame{
			Location: fmt.Sprintf("%s:%d", f, l),
			FuncName: fn.Name(),
		})
	}
	return frames
}

func Traceback() Frames { return traceback(1, 0) }

func CurrentTraceback() *Frame {
	frames := traceback(1, 1)
	if len(frames) == 0 {
		panic("invalid frame list")
	}
	return frames[0]
}
