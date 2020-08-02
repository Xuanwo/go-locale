// +build integration_test

package locale

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDetectViaWin32OLE(t *testing.T) {
	Convey("detect via win32 ole", t, func() {
		langs, err := detectViaWin32OLE()

		Convey("The error should not be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The langs should not be empty", func() {
			So(langs, ShouldNotBeEmpty)
		})
	})
}

func Test_detectViaWinRegistry(t *testing.T) {
	Convey("detect via Windows Registry", t, func() {
		langs, err := detectViaWinRegistry()

		Convey("The error should not be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The langs should not be empty", func() {
			So(langs, ShouldNotBeEmpty)
		})
	})
}
