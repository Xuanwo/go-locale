package locale

import (
	"errors"
	"os"
	"os/exec"

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
	return detect()
}

var lookupEnv = func(env string) (string, bool) {
	return os.LookupEnv(env)
}

var execCommand = func(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.Output()
}
