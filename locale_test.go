package locale

import (
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestDetect(t *testing.T) {
	monkey.Patch(detect, func() (language.Tag, error) {
		return language.English, nil
	})

	tag, err := Detect()
	assert.Nil(t, err)
	assert.Equal(t, language.English, tag)
}
