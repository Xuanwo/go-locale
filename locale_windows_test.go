package locale

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_detectViaRegistry(t *testing.T) {
	Convey("detect via Windows Registry", t, func() {
		langs, err := detectViaRegistry()

		t.Logf("langs: %v", langs)
		Convey("The error should not be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The langs should not be empty", func() {
			So(langs, ShouldNotBeEmpty)
		})
	})
}
