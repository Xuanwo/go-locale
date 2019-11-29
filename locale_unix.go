package locale

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

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
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	if lang != "" {
		tag = language.Make(lang)
		return
	}

	err = ErrNotDetected
	return
}

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

// detectFromEnv will check linux env in order which decided by
// unix convention
//
// ref:
//   - http://man7.org/linux/man-pages/man7/locale.7.html
//   - https://linux.die.net/man/3/gettext
//   - https://wiki.archlinux.org/index.php/Locale
func detectViaEnv() string {
	// Check LANGUAGE: Program use gettext will respect LANGUAGE env
	s, ok := os.LookupEnv("LANGUAGE")
	if ok {
		return parseLanguageEnv(s)[0]
	}

	// Check LC_* in order
	for _, v := range envs {
		s, ok := os.LookupEnv(v)
		if ok {
			return parseLCEnv(s)
		}
	}
	return ""
}

func detectViaLocale() (string, error) {
	errorMessage := "detect via locale: %w"

	cmd := exec.Command("locale")

	var out bytes.Buffer
	cmd.Stdout = &out

	// Output should be like:
	//
	// LANG=en_US.UTF-8
	// LC_CTYPE="en_US.UTF-8"
	// LC_NUMERIC="en_US.UTF-8"
	// LC_TIME="en_US.UTF-8"
	// LC_COLLATE="en_US.UTF-8"
	// LC_MONETARY="en_US.UTF-8"
	// LC_MESSAGES=
	// LC_PAPER="en_US.UTF-8"
	// LC_NAME="en_US.UTF-8"
	// LC_ADDRESS="en_US.UTF-8"
	// LC_TELEPHONE="en_US.UTF-8"
	// LC_MEASUREMENT="en_US.UTF-8"
	// LC_IDENTIFICATION="en_US.UTF-8"
	// LC_ALL=
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf(errorMessage, err)
	}

	m := make(map[string]string)
	s := bufio.NewScanner(&out)
	for s.Scan() {
		value := strings.Split(s.Text(), "=")
		// Ignore not set locale value.
		if len(value) != 2 || value[1] == "" {
			continue
		}
		m[value[0]] = strings.Trim(value[1], "\"")
	}

	for _, v := range envs {
		x, ok := m[v]
		if ok {
			return parseLCEnv(x), nil
		}
	}
	return "", nil
}

// parseLanguageEnv will parse LANGUAGE env.
// Input could be: "en_AU:en_GB:en"
func parseLanguageEnv(s string) []string {
	return strings.Split(s, ":")
}

// parseLCEnv will parse LC_* env.
// Input could be: "en_US.UTF-8"
func parseLCEnv(s string) string {
	x := strings.Split(s, ".")
	// "C" means "ANSI-C" and "POSIX", if locale set to C, we can simple
	// set returned language to "en_US"
	if x[0] == "C" {
		return "en_US"
	}
	return x[0]
}
