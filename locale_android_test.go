//go:build android

package locale

import (
	"testing"
)

func TestDetectViaGetProp(t *testing.T) {
	langs, err := detectViaGetProp()

	t.Logf("langs: %v", langs)
	if err != nil {
		t.Errorf("Expected nil error, got: %v", err)
	}
	if len(langs) == 0 {
		t.Error("Expected non-empty langs, got empty slice")
	}
}
