package main

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	subtests := []struct {
		name                string
		tBotToken, tBTValue string
		tTargetUrl, tTValue string
	}{
		{
			name:      "BOT_TOKEN unset",
			tBotToken: "BOT_TOKEN",
		},
		{
			name:       "TARGET_URL unset",
			tTargetUrl: "TARGET_URL",
		},
		{
			name:       "BOTH set",
			tBotToken:  "BOT_TOKEN",
			tBTValue:   "abc123",
			tTargetUrl: "TARGET_URL",
			tTValue:    "https://google.com",
		},
	}

	// store prev value if set, then unset it
	ptvalBt, ok := os.LookupEnv("BOT_TOKEN")
	if ok && ptvalBt != "" {
		os.Unsetenv("BOT_TOKEN")
	}
	ptvalT, ok := os.LookupEnv("TARGET_URL")
	if ok && ptvalT != "" {
		os.Unsetenv("TARGET_URL")
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			if subtest.tBTValue != "" {
				t.Setenv(subtest.tBotToken, subtest.tBTValue)
			} else {
				err := run()
				require.NotEqual(t, nil, err)
				return
			}
			if subtest.tTValue != "" {
				t.Setenv(subtest.tTargetUrl, subtest.tTValue)
			} else {
				err := run()
				require.NotEqual(t, nil, err)
				return
			}
			go run()

			BotChannel <- syscall.SIGINT
		})
	}

	if ptvalBt != "" {
		os.Setenv("BOT_TOKEN", ptvalBt)
	}
	if ptvalT != "" {
		os.Setenv("TARGET_URL", ptvalT)
	}

	go run()
	BotChannel <- syscall.SIGINT
}
