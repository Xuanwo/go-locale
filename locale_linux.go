package locale

import (
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

func detect() (tag language.Tag, err error) {
	errorMessage := "detect: %w"

	// Check via env firstly.
	lang := detectViaEnv()
	if lang != "" {
		tag = language.Make(lang)
		return
	}

	// Check via locale then.
	lang, err = detectViaLocale()
	if err != nil && !errors.Is(err, ErrNotDetected) {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	if lang != "" {
		tag = language.Make(lang)
		return
	}

	err = fmt.Errorf(errorMessage, ErrNotDetected)
	return
}
