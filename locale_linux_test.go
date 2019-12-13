package locale

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestLinuxDetect(t *testing.T) {
	t.Run("detect via env", func(t *testing.T) {
		monkey.UnpatchAll()

		monkey.Patch(detectViaEnv, func() string {
			return "en_US"
		})

		tag, err := detect()
		assert.Nil(t, err)
		assert.Equal(t, language.AmericanEnglish, tag)
	})

	t.Run("detect via locale with error", func(t *testing.T) {
		monkey.UnpatchAll()

		testError := errors.New("test error")

		monkey.Patch(detectViaEnv, func() string {
			return ""
		})
		monkey.Patch(detectViaLocale, func() (string, error) {
			return "", testError
		})

		_, err := detect()
		assert.True(t, errors.Is(err, testError))
	})

	t.Run("detect via locale", func(t *testing.T) {
		monkey.UnpatchAll()

		monkey.Patch(detectViaEnv, func() string {
			return ""
		})
		monkey.Patch(detectViaLocale, func() (string, error) {
			return "zh_CN", nil
		})

		tag, err := detect()
		assert.Nil(t, err)
		assert.Equal(t, language.Make("zh_CN"), tag)
	})

	t.Run("not detected", func(t *testing.T) {
		monkey.UnpatchAll()

		monkey.Patch(detectViaEnv, func() string {
			return ""
		})
		monkey.Patch(detectViaLocale, func() (string, error) {
			return "", nil
		})

		_, err := detect()
		assert.True(t, errors.Is(err, ErrNotDetected))
	})
}
