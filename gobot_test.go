package main

import (
	"os"
	"testing"
)

func TestMn(t *testing.T) {
	bt, ok := os.LookupEnv("BOT_TOKEN")
	if !ok || bt == "" {
		t.Skip("BOT_TOKEN isn't set")
	}
	t.Setenv("TESTMODE", "true")

	main()
}
