package locale

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestDetectViaEnvLanguage(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     []string
		wantErr  error
	}{
		{"Valid single value", "en_US", []string{"en_US"}, nil},
		{"Multiple values", "en_US:zh_CN", []string{"en_US", "zh_CN"}, nil},
		{"Empty value", "", nil, ErrNotDetected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupEnv()
			defer setupEnv()

			err := os.Setenv("LANGUAGE", tt.envValue)
			if err != nil {
				t.Fatal(err)
			}

			got, err := detectViaEnvLanguage()
			t.Logf("langs: %v", got)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("detectViaEnvLanguage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectViaEnvLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectViaEnvLc(t *testing.T) {
	tests := []struct {
		name    string
		setEnv  bool
		envKey  string
		envVal  string
		want    []string
		wantErr error
	}{
		{"LC_ALL set", true, "LC_ALL", "en_US.UTF-8", []string{"en_US"}, nil},
		{"No LC env set", false, "", "", nil, ErrNotDetected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupEnv()
			defer setupEnv()

			if tt.setEnv {
				err := os.Setenv(tt.envKey, tt.envVal)
				if err != nil {
					t.Fatal(err)
				}
			}

			got, err := detectViaEnvLc()
			t.Logf("langs: %v", got)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("detectViaEnvLc() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectViaEnvLc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseEnvLc(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"en_US.UTF-8", "en_US.UTF-8", "en_US"},
		{"C.UTF-8", "C.UTF-8", "en_US"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseEnvLc(tt.input); got != tt.want {
				t.Errorf("parseEnvLc() = %v, want %v", got, tt.want)
			}
		})
	}
}
