package locale

import (
	"errors"

	"golang.org/x/text/language"
)

var (
	ErrNotDetected = errors.New("not_detected")
)

func Detect() (tag language.Tag, err error) {
	return detect()
}
