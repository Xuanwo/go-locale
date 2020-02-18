package locale

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func hookLookupEnv(hook func(string) (string, bool)) func() {
	old := lookupEnv
	lookupEnv = hook
	return func() { lookupEnv = old }
}

func hookExecCommand(hook func(string, ...string) ([]byte, error)) func() {
	old := execCommand
	execCommand = hook
	return func() { execCommand = old }
}

func TestDetect(t *testing.T) {
	old := detect
	detect = func() (language.Tag, error) {
		return language.English, nil
	}

	tag, err := Detect()
	assert.Nil(t, err)
	assert.Equal(t, language.English, tag)

	detect = old
}
