package main

import (
	"os"
	"syscall"
	"testing"
)

func TestFuncMain(t *testing.T) {
	bt, ok := os.LookupEnv("BOT_TOKEN")
	if !ok || bt == "" {
		t.Skip("BOT_TOKEN isn't set")
	}

	go main()
	BotChannel <- syscall.SIGINT
}
