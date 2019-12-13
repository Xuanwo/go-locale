package locale

import (
	"testing"
)

func TestDetectViaUserDefaultsSystem(t *testing.T) {
	v, err := detectViaUserDefaultsSystem()
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
