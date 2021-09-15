package failure

import "testing"

func TestError(t *testing.T) {
	e := New(Errno(0), "success")
	t.Logf("%+v\n", e)
	t.Logf("%+v\n", e.frame)
}
