//go:build !windows && !darwin && !js
// +build !windows,!darwin,!js

package locale

import (
	"os"
	"path"
	"testing"
)

func TestXDGConfigHome(t *testing.T) {
	setupEnv()
	tmpDir := setupLocaleConf("locale.conf")
	defer os.RemoveAll(tmpDir)

	err := os.Setenv("XDG_CONFIG_HOME", tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	fp := getLocaleConfPath()
	expected := path.Join(tmpDir, "locale.conf")
	if fp != expected {
		t.Errorf("Expected path %s, got %s", expected, fp)
	}
}

func TestHOME(t *testing.T) {
	setupEnv()
	tmpDir := setupLocaleConf(".config/locale.conf")
	defer os.RemoveAll(tmpDir)

	err := os.Setenv("HOME", tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	fp := getLocaleConfPath()
	expected := path.Join(tmpDir, ".config/locale.conf")
	if fp != expected {
		t.Errorf("Expected path %s, got %s", expected, fp)
	}
}

func TestFallbackToSystem(t *testing.T) {
	setupEnv()
	var localeExist bool
	_, err := os.Stat("/etc/locale.conf")
	if err == nil {
		localeExist = true
	}

	fp := getLocaleConfPath()
	if (fp == "/etc/locale.conf") != localeExist {
		t.Errorf("Expected path to be /etc/locale.conf: %v, got: %s", localeExist, fp)
	}
}

func TestDetectViaLocaleConf(t *testing.T) {
	setupEnv()
	tmpDir := setupLocaleConf("locale.conf")
	defer os.RemoveAll(tmpDir)

	err := os.Setenv("XDG_CONFIG_HOME", tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	lang, err := detectViaLocaleConf()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(lang) == 0 {
		t.Error("Expected non-empty lang, got empty string")
	}
}
