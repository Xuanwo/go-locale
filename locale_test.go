// +build unit_test

package locale

import (
	"errors"
	"testing"

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
