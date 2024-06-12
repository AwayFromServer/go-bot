package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdStatus(t *testing.T) {
	testdata := []struct {
		name         string
		target       string
		expected     string
		expected_err string
	}{
		{
			"Valid status check",
			"https://dns.google",
			"It looks like it's up!",
			"",
		},
		{
			"Failed status check",
			"this is not a website",
			"It looks like it's offline...",
			"Get \"this%20is%20not%20a%20website\": unsupported protocol scheme \"\"",
		},
	}

	for _, subtest := range testdata {
		t.Run(subtest.name, func(t *testing.T) {
			actual, err := getCurrentStatus(subtest.target)

			assert.Equal(t, subtest.expected, actual.Content)
			if err != nil {
				assert.Equal(t, subtest.expected_err, err.Error())
			}

		})
	}
	target := "https://google.com"
	expected := "It looks like it's up!"
	actual, err := getCurrentStatus(target)

	assert.Equal(t, expected, actual.Content)
	assert.Equal(t, nil, err)

	target = "this is not a website"
	expected = "It looks like it's offline..."
	actual, err = getCurrentStatus(target)

	assert.Equal(t, expected, actual.Content)
	assert.NotEqual(t, nil, err)
}
