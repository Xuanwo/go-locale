//go:build android

package locale

import (
	"bytes"
	"os/exec"
	"strings"
)

var detectors = []detector{
	detectViaEnvLanguage,
	detectViaEnvLc,
	detectViaGetProp,
}

var androidGetPropKeys = []string{
	"persist.sys.locale",
	"ro.product.locale",
}

func detectViaGetProp() ([]string, error) {
	for _, key := range androidGetPropKeys {
		lang, err := getSystemProperty(key)
		if err == nil {
			return []string{lang}, nil
		}
	}
	return nil, &Error{"detect via getprop", ErrNotDetected}
}

func getSystemProperty(key string) (string, error) {
	cmd := exec.Command("getprop", key)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", &Error{"detect via getprop", err}
	}

	content := strings.TrimSpace(out.String())
	if content == "" {
		return "", &Error{"detect via getprop", ErrNotDetected}
	}
	return content, nil
}
