package locale

import (
	"errors"
	"os"
	"strings"

	"golang.org/x/text/language"
)

var (
	// ErrNotDetected returns while no locale detected.
	ErrNotDetected = errors.New("not detected")
	// ErrNotSupported means current platform or language is not supported.
	ErrNotSupported = errors.New("not supported")
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

	tags = make([]language.Tag, 0, len(lang))
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
	return nil, &Error{"detect", ErrNotDetected}
}

type detector func() ([]string, error)

// Unless we call LookupEnv more than 9 times, we should not use Environ.
//
// goos: linux
// goarch: amd64
// pkg: github.com/Xuanwo/go-locale
// BenchmarkLookupEnv
// BenchmarkLookupEnv-8   	37024654	        32.4 ns/op
// BenchmarkEnviron
// BenchmarkEnviron-8     	 4275735	       281 ns/op
// PASS

// envs is the env to be checked.
//
// LC_ALL will overwrite all LC_* options.
// FIXME: LC_ALL=C should overwrite $LANGUAGE env
//
// LC_MESSAGES is the config for messages.
// FIXME: LC_MESSAGES=C should overwrite $LANGUAGE env
//
// LANG is the default locale.
var envs = []string{"LC_ALL", "LC_MESSAGES", "LANG"}

// detectViaEnvLanguage checks env LANGUAGE
//
// Program use gettext will respect LANGUAGE env
func detectViaEnvLanguage() ([]string, error) {
	s, ok := os.LookupEnv("LANGUAGE")
	if !ok || s == "" {
		return nil, &Error{"detect via env language", ErrNotDetected}
	}
	return parseEnvLanguage(s), nil
}

// detectViaEnvLc checks LC_* in order which decided by
// unix convention
//
// ref:
//   - http://man7.org/linux/man-pages/man7/locale.7.html
//   - https://linux.die.net/man/3/gettext
//   - https://wiki.archlinux.org/index.php/Locale
func detectViaEnvLc() ([]string, error) {
	for _, v := range envs {
		s, ok := os.LookupEnv(v)
		if ok && s != "" {
			return []string{parseEnvLc(s)}, nil
		}
	}
	return nil, &Error{"detect via env lc", ErrNotDetected}
}

// parseEnvLanguage will parse LANGUAGE env.
// Input could be: "en_AU:en_GB:en"
func parseEnvLanguage(s string) []string {
	return strings.Split(s, ":")
}

// parseEnvLc will parse LC_* env.
// Input could be: "en_US.UTF-8"
func parseEnvLc(s string) string {
	x := strings.Split(s, ".")
	// "C" means "ANSI-C" and "POSIX", if locale set to C, we can simple
	// set returned language to "en_US"
	if x[0] == "C" {
		return "en_US"
	}
	return x[0]
}
