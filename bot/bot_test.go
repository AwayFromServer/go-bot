package bot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	expected_token := "abc123"
	expected_url := "https://google.com"
	t.Setenv("BOT_TOKEN", expected_token)
	t.Setenv("TARGET_URL", expected_url)

	b := New(expected_token, expected_url)

	assert.Equal(t, b.botToken, expected_token)
}

func TestRun(t *testing.T) {
	expected_token := "abc123"
	expected_url := "https://google.com"
	t.Setenv("TARGET_URL", expected_url)

	b := New(expected_token, expected_url)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go b.Run(ctx)

	if b.targetURL != expected_url { // placeholder -> bot functions get tested here
		t.Errorf("TARGET_URL = %v, want %v", b.targetURL, expected_url)
	}
}
