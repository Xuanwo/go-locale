// +build !tests

package locale

import (
	"testing"


)

func TestDetectViaUserDefaultsSystem(t *testing.T) {
	Convey("detect via user defaults system", t, func() {
		langs, err := detectViaUserDefaultsSystem()

		Convey("The error should not be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The langs should not be empty", func() {
			So(langs, ShouldNotBeEmpty)
		})
	})
}
