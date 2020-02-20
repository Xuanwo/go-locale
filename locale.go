package locale

import (
	"errors"

	"golang.org/x/text/language"
)

var (
	// ErrNotDetected returns while no locale detected.
	ErrNotDetected = errors.New("not detected")
	// ErrSystemError returns while error happened in system call.
	ErrSystemError = errors.New("system error")
)

// Detect will detect current env's language.
func Detect() (tag language.Tag, err error) {
	tags, err := DetectAll()
	if err != nil {
		return language.Und, err
	}
	return tags[0], nil
}

// DetectAll will detect current env's all available language.
func DetectAll() (tags []language.Tag, err error) {
	lang, err := detect()
	if err != nil {
		return
	}

	tags = make([]language.Tag, 0,len(lang))
	for _, v := range lang {
		tags = append(tags, language.Make(v))
	}
	return
}

func detect() (lang []string, err error) {
	for _, fn := range detectors {
		lang, err = fn()
		if err != nil && errors.Is(err, ErrNotDetected) {
			continue
		}
		if err != nil {
			return
		}
		return
	}
	return nil, ErrNotDetected
}

type detector func() ([]string, error)
