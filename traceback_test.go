package failure

import "testing"

func TestTraceback(t *testing.T) {
	frames := b()
	for _, f := range frames {
		t.Logf("frames: %+v", f)
	}
}

func a() Frames {
	return Traceback()
}

func b() Frames {
	return a()
}
