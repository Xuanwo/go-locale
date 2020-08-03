// +build unit_test

package locale

import (
	"errors"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestInternalDetect(t *testing.T) {
	teseerr := errors.New("test error")
	tests := []struct {
		name         string
		mockString   []string
		mockError    error
		expectString []string
		expectError  error
	}{
		{
			"normal",
			[]string{"en_US"},
			nil,
			[]string{"en_US"},
			nil,
		},
		{
			"not detected",
			[]string(nil),
			ErrNotDetected,
			[]string(nil),
			ErrNotDetected,
		},
		{
			"not detected",
			[]string(nil),
			teseerr,
			[]string(nil),
			teseerr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLang.set(tt.mockString, tt.mockError)

			lang, err := detect()
			if tt.expectError != nil {
				assert.True(t, errors.Is(err, tt.expectError))
				assert.Empty(t, lang)
			} else {
				assert.Nil(t, err)
				assert.EqualValues(t, tt.expectString, lang)
			}
		})
	}
}

func TestDetect(t *testing.T) {
	tests := []struct {
		name        string
		mockString  []string
		mockError   error
		expectLang  language.Tag
		expectError error
	}{
		{
			"normal",
			[]string{"en-US"},
			nil,
			language.AmericanEnglish,
			nil,
		},
		{
			"not detected",
			[]string(nil),
			ErrNotDetected,
			language.Und,
			ErrNotDetected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLang.set(tt.mockString, tt.mockError)

			lang, err := Detect()
			if tt.expectError != nil {
				assert.True(t, errors.Is(err, tt.expectError))
				assert.Empty(t, lang)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectLang, lang)
			}
		})
	}
}

func TestDetectAll(t *testing.T) {
	tests := []struct {
		name        string
		mockString  []string
		mockError   error
		expectLang  []language.Tag
		expectError error
	}{
		{
			"normal",
			[]string{"en-US"},
			nil,
			[]language.Tag{language.AmericanEnglish},
			nil,
		},
		{
			"not detected",
			[]string(nil),
			ErrNotDetected,
			[]language.Tag(nil),
			ErrNotDetected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLang.set(tt.mockString, tt.mockError)

			lang, err := DetectAll()
			if tt.expectError != nil {
				assert.True(t, errors.Is(err, tt.expectError))
				assert.Empty(t, lang)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectLang, lang)
			}
		})
	}
}

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
