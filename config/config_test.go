package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	testData := []struct {
		name        string
		file        string
		expected    []string
		expectedErr string
	}{
		{
			"Testing JSON",
			"config_test.json",
			[]string{"foo", "bar"},
			"",
		},
		{
			"Testing YAML",
			"config_test.yaml",
			[]string{"baz", "qux"},
			"",
		},
		{
			"Testing NoFile Error",
			"notarealfile.json",
			[]string{""},
			"open notarealfile.json: no such file or directory",
		},
		{
			"Testing Unmarshal Error",
			"bork_test.json",
			[]string{""},
			"invalid character 'v' after object key",
		},
	}

	for _, subtest := range testData {
		t.Run(subtest.name, func(t *testing.T) {
			actual, err := GetConfig(subtest.file)

			if err != nil {
				assert.Equal(t, subtest.expectedErr, err.Error())
			} else {
				assert.Equal(t, subtest.expected[0], actual[0]["botPrefix"])
				assert.Equal(t, subtest.expected[1], actual[0]["botToken"])
			}
		})
	}
}

func TestGetOverrides(t *testing.T) {
	testData := []struct {
		name           string
		expectedPrefix string
		expectedToken  string
		overridePrefix string
		overrideToken  string
	}{
		{
			"Testing No Overrides",
			"foo",
			"bar",
			"",
			"",
		},
		{
			"Testing Override BP",
			"pip",
			"bar",
			"pip",
			"",
		},
		{
			"Testing Override BT",
			"foo",
			"pop",
			"",
			"pop",
		},
		{
			"Testing DoubleOverrides",
			"pip",
			"pop",
			"pip",
			"pop",
		},
	}

	for _, subtest := range testData {
		t.Run(subtest.name, func(t *testing.T) {
			t.Setenv("botPrefix", subtest.overridePrefix)
			t.Setenv("botToken", subtest.overrideToken)
			actual, err := GetConfig("config_test.json")
			if err != nil {
				t.Fatal(err)
			}
			actual = GetOverrides(actual)
			assert.Equal(t, subtest.expectedPrefix, actual[0]["botPrefix"])
			assert.Equal(t, subtest.expectedToken, actual[0]["botToken"])
		})
	}
}
