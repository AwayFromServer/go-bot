package bot

import (
	"testing"
)

func TestSetup(t *testing.T) {
	expected_token := "abc123"
	expected_url := "https://google.com"
	t.Setenv("BOT_TOKEN", expected_token)
	t.Setenv("HEARTBEAT_URL", expected_url)
	Setup()

	switch {
	case botToken != expected_token:
		if botToken == "" {
			t.Fatal(`botToken is unset`)
		}
		t.Fatalf(`BOT_TOKEN = %q, want %v`, botToken, expected_token)
	case heartbeatURL != expected_url:
		if heartbeatURL == "" {
			t.Fatal(`heartbeatURL is unset`)
		}
		t.Fatalf(`HEARTBEAT_URL = %q, want %v`, heartbeatURL, expected_url)
	}
}

func TestRun(t *testing.T) {

}
