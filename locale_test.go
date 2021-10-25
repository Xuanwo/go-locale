package locale

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

var mockLang mock

type mock struct {
	s   []string
	err error

	sync.Mutex
}

func (l *mock) get() ([]string, error) {
	l.Lock()
	defer l.Unlock()

	return l.s, l.err
}

func (l *mock) set(s []string, e error) {
	l.Lock()
	defer l.Unlock()

	l.s = s
	l.err = e
}

func TestInternalDetect(t *testing.T) {
	detectors = []detector{mockLang.get}

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
	detectors = []detector{mockLang.get}

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
	detectors = []detector{mockLang.get}

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
