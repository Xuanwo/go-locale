// +build darwin dragonfly freebsd linux netbsd openbsd solaris illumos
// +build integration_test

package locale

import (
	"errors"
	"os"
	"strings"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDetectViaEnvLanguage(t *testing.T) {
	Convey("detect via env language", t, func() {
		// Make sure env has clear before current test.
		setupEnv()

		Reset(func() {
			// Reset all env after every Convey.
			setupEnv()
		})

		Convey("When LANGUAGE has valid value", func() {
			err := os.Setenv("LANGUAGE", "en_US")
			if err != nil {
				t.Error(err)
			}

			lang, err := detectViaEnvLanguage()

			Convey("The error should not be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The lang should not be equal", func() {
				So(lang, ShouldResemble, []string{"en_US"})
			})
		})

		Convey("When LANGUAGE has multiple value", func() {
			err := os.Setenv("LANGUAGE", "en_US:zh_CN")
			if err != nil {
				t.Error(err)
			}

			lang, err := detectViaEnvLanguage()

			Convey("The error should not be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The lang should not be equal", func() {
				So(lang, ShouldResemble, []string{"en_US", "zh_CN"})
			})
		})

		Convey("When LANGUAGE is empty", func() {
			err := os.Setenv("LANGUAGE", "")
			if err != nil {
				t.Error(err)
			}

			lang, err := detectViaEnvLanguage()

			Convey("The error should be ErrNotDetected", func() {
				So(errors.Is(err, ErrNotDetected), ShouldBeTrue)
			})
			Convey("The lang should be empty", func() {
				So(lang, ShouldBeEmpty)
			})
		})
	})
}

func TestDetectViaEnvLc(t *testing.T) {
	Convey("detect via env language", t, func() {
		// Make sure env has clear before current test.
		setupEnv()

		Reset(func() {
			// Reset all env after every Convey.
			setupEnv()
		})

		Convey("When LC_ALL has been set", func() {
			err := os.Setenv("LC_ALL", "en_US.UTF-8")
			if err != nil {
				t.Error(err)
			}

			lang, err := detectViaEnvLc()

			Convey("The error should not be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The lang should not be equal", func() {
				So(lang, ShouldResemble, []string{"en_US"})
			})
		})

		Convey("When no LC env has been set", func() {
			lang, err := detectViaEnvLc()

			Convey("The error should be ErrNotDetected", func() {
				So(errors.Is(err, ErrNotDetected), ShouldBeTrue)
			})
			Convey("The lang should be empty", func() {
				So(lang, ShouldBeEmpty)
			})
		})
	})
}

func TestDetectViaLocale(t *testing.T) {
	Convey("detect via locale", t, func() {
		lang, err := detectViaLocale()

		Convey("The error should not be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The lang should not be empty", func() {
			So(lang, ShouldNotBeEmpty)
		})
	})
}

var env struct {
	Env map[string]string
	sync.Mutex
	sync.Once
}

func setupEnv() {
	env.Lock()
	defer env.Unlock()

	env.Do(func() {
		env.Env = make(map[string]string)
		for _, v := range os.Environ() {
			x := strings.SplitN(v, "=", 2)
			// Ignore all language related env
			if strings.HasPrefix(x[0], "LANG") || strings.HasPrefix(x[0], "LC") {
				continue
			}
			env.Env[x[0]] = x[1]
		}
	})

	os.Clearenv()

	for k, v := range env.Env {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
	return
}
