// +build !windows !js
// +build !integration_test

package locale

import (
	"os"
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetLocaleConfPath(t *testing.T) {
	Convey("get locale conf path", t, func() {
		// Make sure env has clear before current test.
		setupEnv()

		Reset(func() {
			// Reset all env after every Convey.
			setupEnv()
		})

		Convey("When user set XDG_CONFIG_HOME", func() {
			tmpDir := setupLocaleConf("locale.conf")
			Reset(func() {
				_ = os.RemoveAll(tmpDir)
			})

			err := os.Setenv("XDG_CONFIG_HOME", tmpDir)
			if err != nil {
				t.Error(err)
			}

			fp := getLocaleConfPath()

			Convey("The path should be equal", func() {
				So(fp, ShouldEqual, path.Join(tmpDir, "locale.conf"))
			})
		})

		Convey("When user set HOME", func() {
			tmpDir := setupLocaleConf(".config/locale.conf")
			Reset(func() {
				_ = os.RemoveAll(tmpDir)
			})

			err := os.Setenv("HOME", tmpDir)
			if err != nil {
				t.Error(err)
			}

			fp := getLocaleConfPath()

			Convey("The path should be equal", func() {
				So(fp, ShouldEqual, path.Join(tmpDir, ".config/locale.conf"))
			})
		})

		Convey("When fallback to system level locale.conf", func() {
			var localeExist bool
			_, err := os.Stat("/etc/locale.conf")
			if err == nil {
				localeExist = true
			}

			fp := getLocaleConfPath()

			Convey("The path should be equal", func() {
				So(fp == "/etc/locale.conf", ShouldEqual, localeExist)
			})
		})
	})
}

func TestDetectViaLocaleConf(t *testing.T) {
	Convey("detect via locale conf", t, func() {
		setupEnv()
		Reset(func() {
			setupEnv()
		})

		tmpDir := setupLocaleConf("locale.conf")
		Reset(func() {
			_ = os.RemoveAll(tmpDir)
		})
		err := os.Setenv("XDG_CONFIG_HOME", tmpDir)
		if err != nil {
			t.Error(err)
		}

		lang, err := detectViaLocaleConf()

		Convey("The error should be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The lang should not be empty", func() {
			So(lang, ShouldNotBeEmpty)
		})
	})
}
