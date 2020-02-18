package locale

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func hookDetects(Env func() string, Locale func() (string, error)) func() {
	oldEnv := detectViaEnv
	oldLocale := detectViaLocale

	if Env != nil {
		detectViaEnv = Env
	}
	if Locale != nil {
		detectViaLocale = Locale
	}

	return func() {
		detectViaEnv = oldEnv
		detectViaLocale = oldLocale
	}
}

func TestLinuxDetect(t *testing.T) {
	t.Run("detect via env", func(t *testing.T) {
		unhook := hookDetects(func() string { return "en_US" }, nil)
		defer unhook()

		tag, err := detect()
		assert.Nil(t, err)
		assert.Equal(t, language.AmericanEnglish, tag)
	})

	t.Run("detect via locale with error", func(t *testing.T) {
		testError := errors.New("test error")

		unhook := hookDetects(
			func() string { return "" },
			func() (string, error) { return "", testError },
		)
		defer unhook()

		_, err := detect()
		assert.True(t, errors.Is(err, testError))
	})

	t.Run("detect via locale", func(t *testing.T) {
		unhook := hookDetects(
			func() string { return "" },
			func() (string, error) { return "zh_CN", nil },
		)
		defer unhook()

		tag, err := detect()
		assert.Nil(t, err)
		assert.Equal(t, language.Make("zh_CN"), tag)
	})

	t.Run("not detected", func(t *testing.T) {
		unhook := hookDetects(
			func() string { return "" },
			func() (string, error) { return "", nil },
		)
		defer unhook()

		_, err := detect()
		assert.True(t, errors.Is(err, ErrNotDetected))
	})
}
