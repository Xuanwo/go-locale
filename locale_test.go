package locale

import (
	"errors"
	"reflect"
	"sync"
	"testing"

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

	testErr := errors.New("test error")
	tests := []struct {
		name         string
		mockString   []string
		mockError    error
		expectString []string
		expectError  error
	}{
		{"normal", []string{"en_US"}, nil, []string{"en_US"}, nil},
		{"not detected", nil, ErrNotDetected, nil, ErrNotDetected},
		{"other error", nil, testErr, nil, testErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLang.set(tt.mockString, tt.mockError)

			lang, err := detect()
			if !errors.Is(err, tt.expectError) {
				t.Errorf("detect() error = %v, expectError %v", err, tt.expectError)
			}
			if !reflect.DeepEqual(lang, tt.expectString) {
				t.Errorf("detect() = %v, want %v", lang, tt.expectString)
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
		{"normal", []string{"en-US"}, nil, language.AmericanEnglish, nil},
		{"invalid", []string{"ac"}, nil, language.Und, nil},
		{"not detected", nil, ErrNotDetected, language.Und, ErrNotDetected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLang.set(tt.mockString, tt.mockError)

			lang, err := Detect()
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Detect() error = %v, expectError %v", err, tt.expectError)
			}
			if lang != tt.expectLang {
				t.Errorf("Detect() = %v, want %v", lang, tt.expectLang)
			}
		})
	}
}

func BenchmarkDetect(b *testing.B) {
	detectors = []detector{mockLang.get}

	mockLang.set([]string{"en-US"}, nil)
	for i := 0; i < b.N; i++ {
		Detect()
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
		{"normal", []string{"en-US"}, nil, []language.Tag{language.AmericanEnglish}, nil},
		{"not detected", nil, ErrNotDetected, nil, ErrNotDetected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLang.set(tt.mockString, tt.mockError)

			lang, err := DetectAll()
			if !errors.Is(err, tt.expectError) {
				t.Errorf("DetectAll() error = %v, expectError %v", err, tt.expectError)
			}
			if !reflect.DeepEqual(lang, tt.expectLang) {
				t.Errorf("DetectAll() = %v, want %v", lang, tt.expectLang)
			}
		})
	}
}
