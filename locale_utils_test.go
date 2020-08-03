package locale

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"testing"
	"time"
)

var env struct {
	Env map[string]string

	sync.Mutex
	sync.Once
}

func setupLocaleConf(filePath string) (dir string) {
	confContent := `LANG=en_US.UTF-8`
	tmpDir := "/tmp/" + time.Now().String()
	baseDir := path.Dir(path.Join(tmpDir, filePath))

	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path.Join(tmpDir, filePath), []byte(confContent), 0644)
	if err != nil {
		panic(err)
	}

	return tmpDir
}

func setupEnv() {
	env.Lock()
	defer env.Unlock()

	env.Do(func() {
		env.Env = make(map[string]string)
		for _, v := range os.Environ() {
			x := strings.SplitN(v, "=", 2)
			// Ignore all language related env
			if strings.HasPrefix(x[0], "LANG") || strings.HasPrefix(x[0], "LC") {
				continue
			}
			env.Env[x[0]] = x[1]
		}
	})

	os.Clearenv()

	for k, v := range env.Env {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
	return
}

func BenchmarkLookupEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = os.LookupEnv("LANGUAGE")
	}
}

func BenchmarkEnviron(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = os.Environ()
	}
}
