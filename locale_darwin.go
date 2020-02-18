package locale

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

var detect = func() (tag language.Tag, err error) {
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

	lang, err = detectViaUserDefaultsSystem()
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

// detectViaUserDefaultsSystem will detect language via Apple User Defaults System
//
// ref: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/UserDefaults/AboutPreferenceDomains/AboutPreferenceDomains.html
func detectViaUserDefaultsSystem() (string, error) {
	errorMessage := "detect via defaults: %w"

	out, err := execCommand("defaults", "read", "NSGlobalDomain", "AppleLanguages")
	if err != nil {
		return "", fmt.Errorf(errorMessage, err)
	}

	// Output should be like:
	//
	// (
	//    en,
	//    ja,
	//    fr,
	//    de,
	//    es,
	//    it,
	//    pt,
	//    "pt-PT",
	//    nl,
	//    sv,
	//    nb,
	//    da,
	//    fi,
	//    ru,
	//    pl,
	//    "zh-Hans",
	//    "zh-Hant",
	//    ko,
	//    ar,
	//    cs,
	//    hu,
	//    tr
	// )

	lines := strings.Split(string(out), "\n")
	m := make([]string, 0, len(lines))

	for _, text := range lines {
		// Ignore "(" and ")"
		if !strings.HasPrefix(text, " ") {
			continue
		}
		// Trim all space, " and ,
		text = strings.Trim(text, " \",")
		m = append(m, text)
	}

	if len(m) > 0 {
		return m[0], nil
	}
	return "", nil
}
