//go:build android

package locale

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var detectors = []detector{
	detectViaEnvLanguage,
	detectViaEnvLc,
	detectViaGetProp,
}

var androidLocaleKeys = []string{
	"persist.sys.locale",
	"ro.product.locale",
	"persist.sys.language",
}

var androidGetPropPaths = []string{
	"/system/bin/getprop",
	"getprop",
}

func detectViaGetProp() ([]string, error) {
	for _, key := range androidLocaleKeys {
		lang, err := getSystemProperty(key)
		if err == nil {
			return []string{lang}, nil
		}
	}
	lang, country := tryCombinedLocale()
	if lang != "" && country != "" {
		return []string{fmt.Sprintf("%s-%s", lang, country)}, nil
	}
	lang, country = tryCombinedLocaleAlt()
	if lang != "" && country != "" {
		return []string{fmt.Sprintf("%s-%s", lang, country)}, nil
	}
	return nil, &Error{"detect via getprop", ErrNotDetected}
}

func tryCombinedLocale() (string, string) {
	lang, err := getSystemProperty("persist.sys.language")
	if err != nil {
		return "", ""
	}
	country, err := getSystemProperty("persist.sys.country")
	if err != nil {
		return "", ""
	}
	if lang == "" || country == "" {
		return "", ""
	}
	return lang, country
}

func tryCombinedLocaleAlt() (string, string) {
	lang, err := getSystemProperty("ro.product.locale.language")
	if err != nil {
		return "", ""
	}
	country, err := getSystemProperty("ro.product.locale.region")
	if err != nil {
		return "", ""
	}
	return lang, country
}

func getSystemProperty(key string) (string, error) {
	for _, path := range androidGetPropPaths {
		cmd := exec.Command(path, key)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			continue
		}
		content := strings.TrimSpace(out.String())
		if content == "" {
			continue
		}
		return content, nil
	}
	return "", &Error{"detect via getprop", ErrNotDetected}
}
