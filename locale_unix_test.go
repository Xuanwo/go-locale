// +build darwin dragonfly freebsd linux netbsd openbsd

package locale

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectViaEnv(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(key string) (string, bool)
		expected string
	}{
		{
			"env LANGUAGE",
			func(key string) (s string, b bool) {
				assert.Equal(t, "LANGUAGE", key)
				return "en_US", true
			},
			"en_US",
		},
		{
			"env LANG",
			func(key string) (s string, b bool) {
				if key == "LANGUAGE" {
					return "", false
				}
				assert.Contains(t, envs, key)
				if key != "LANG" {
					return "", false
				}
				return "en_US", true
			},
			"en_US",
		},
		{
			"env LC_ALL=C",
			func(key string) (s string, b bool) {
				if key == "LANGUAGE" {
					return "", false
				}
				assert.Contains(t, envs, key)
				return "C", true
			},
			"en_US",
		},
		{
			"no nev",
			func(key string) (s string, b bool) {
				return "", false
			},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unhook := hookLookupEnv(tt.fn)
			defer unhook()
			assert.Equal(t, tt.expected, detectViaEnv())
		})
	}

}

func TestDetectViaLocale(t *testing.T) {
	testError := errors.New("test error")

	tests := []struct {
		name          string
		fn            func(name string, args ...string) ([]byte, error)
		expected      string
		expectedError error
	}{
		{
			"normal case",
			func(name string, args ...string) ([]byte, error) {
				return []byte(`LANG=en_US.UTF-8
LC_CTYPE="en_US.UTF-8"
LC_NUMERIC="en_US.UTF-8"
LC_TIME="en_US.UTF-8"
LC_COLLATE="en_US.UTF-8"
LC_MONETARY="en_US.UTF-8"
LC_MESSAGES=
LC_PAPER="en_US.UTF-8"
LC_NAME="en_US.UTF-8"
LC_ADDRESS="en_US.UTF-8"
LC_TELEPHONE="en_US.UTF-8"
LC_MEASUREMENT="en_US.UTF-8"
LC_IDENTIFICATION="en_US.UTF-8"
LC_ALL=`), nil
			},
			"en_US",
			nil,
		},
		{
			"locale returns error",
			func(name string, args ...string) ([]byte, error) {
				return nil, testError
			},
			"",
			testError,
		},
		{
			"locale returns nothing",
			func(name string, args ...string) ([]byte, error) {
				return []byte(`LANG=
LC_CTYPE=
LC_NUMERIC=
LC_TIME=
LC_COLLATE=
LC_MONETARY=
LC_MESSAGES=
LC_PAPER=
LC_NAME=
LC_ADDRESS=
LC_TELEPHONE=
LC_MEASUREMENT=
LC_IDENTIFICATION=
LC_ALL=`), nil
			},
			"",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unhook := hookExecCommand(tt.fn)
			defer unhook()

			got, err := detectViaLocale()
			assert.True(t, errors.Is(err, tt.expectedError))
			assert.Equal(t, tt.expected, got)
		})
	}
}
