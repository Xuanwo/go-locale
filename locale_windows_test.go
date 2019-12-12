package locale

import (
	"testing"
)

func TestDetectViaWin32OLE(t *testing.T) {
	v, err := detectViaWin32OLE()
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
