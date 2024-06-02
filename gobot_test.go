package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncMain(t *testing.T) {
	expected := "abc123"

	t.Setenv("BOT_TOKEN", expected)

	actual := getBotToken()

	assert.Equal(t, expected, actual)
}
