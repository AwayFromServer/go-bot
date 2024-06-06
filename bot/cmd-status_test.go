package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdStatus(t *testing.T) {
	target := "https://google.com"
	expected := "It looks like it's up! " + target
	actual, err := getCurrentStatus(target)

	assert.Equal(t, expected, actual.Content)
	assert.Equal(t, nil, err)

	expected = "It looks like it's offline... " + target
	actual, err = getCurrentStatus("this is not a website")

	assert.Equal(t, expected, actual.Content)
	assert.NotEqual(t, nil, err)
}
