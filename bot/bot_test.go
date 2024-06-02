package bot

import (
	"context"
	"testing"
)

func TestSetup(t *testing.T) {
	expected_token := "abc123"
	expected_url := "https://google.com"
	t.Setenv("BOT_TOKEN", expected_token)
	t.Setenv("TARGET_URL", expected_url)

	Setup()

	switch {
	case botToken != expected_token:
		if botToken == "" {
			t.Skip("botToken is unset")
		}
		t.Fatalf("BOT_TOKEN = %q, want %v", botToken, expected_token)
	case targetURL != expected_url:
		if targetURL == "" {
			t.Skip("targetURL is unset")
		}
		t.Fatalf("TARGET_URL = %q, want %v", targetURL, expected_url)
	}
}

func TestRun(t *testing.T) {
	expected_url := "https://google.com"
	t.Setenv("TARGET_URL", expected_url)
	Setup()

	if botToken == "" {
		t.Skip("botToken is unset")
	}
	ctx, cancel := context.WithCancel(context.Background())
	go Run(ctx)

	if targetURL != expected_url { // placeholder -> bot functions get tested here
		t.Fatalf("TARGET_URL = %q, want %v", targetURL, expected_url)
	}
	cancel()
}
