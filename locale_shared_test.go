// +build !integration_test

package locale

import (
	"errors"
	"os"
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

func TestParseEnvLc(t *testing.T) {
	Convey("parse env lc", t, func() {
		Convey("When input en_US.UTF-8", func() {
			x := parseEnvLc("en_US.UTF-8")

			Convey("The lang should be en_US", func() {
				So(x, ShouldEqual, "en_US")
			})
		})

		Convey("When input C.UTF-8", func() {
			x := parseEnvLc("C.UTF-8")

			Convey("The lang should be en_US", func() {
				So(x, ShouldEqual, "en_US")
			})
		})
	})
}
