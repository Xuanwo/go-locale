// +build darwin dragonfly freebsd linux netbsd openbsd solaris illumos
// +build !unit_test

package locale

import (
	"bufio"
	"os"
	"path"
	"strings"
)

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

func detectViaLocaleConf() (_ []string, err error) {
	defer func() {
		if err != nil {
			err = &Error{"detect via locale conf", err}
		}
	}()

	fp := getLocaleConfPath()
	if fp == "" {
		return nil, ErrNotDetected
	}

	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

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
	m := make(map[string]string)
	s := bufio.NewScanner(f)
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
			return []string{parseEnvLc(x)}, nil
		}
	}
	return nil, ErrNotDetected
}

// getLocaleConfPath will try to get correct locale conf path.
//
// Following path could be returned:
//   - "$XDG_CONFIG_HOME/locale.conf" (follow XDG Base Directory specification)
//   - "$HOME/.config/locale.conf" (user level locale config)
//   - "/etc/locale.conf" (system level locale config)
//   - "" (empty means no valid path found, caller need to handle this.)
//
// ref:
//   - POSIX Locale: https://pubs.opengroup.org/onlinepubs/9699919799/
//   - XDG Base Directory: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func getLocaleConfPath() string {
	// Try to loading from $XDG_CONFIG_HOME/locale.conf
	xdg, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		fp := path.Join(xdg, "locale.conf")
		_, err := os.Stat(fp)
		if err == nil {
			return fp
		}
	}

	// Try to loading from $HOME/.config/locale.conf
	home, ok := os.LookupEnv("HOME")
	if ok {
		fp := path.Join(home, ".config", "locale.conf")
		_, err := os.Stat(fp)
		if err == nil {
			return fp
		}
	}

	// Try to loading from /etc/locale.conf
	fp := "/etc/locale.conf"
	_, err := os.Stat(fp)
	if err == nil {
		return fp
	}

	return ""
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
