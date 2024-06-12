package main

import (
	"os"

	"github.com/awayfromserver/gobot/bot"
)

func main() {
	if bot.Run() != nil {
		os.Exit(1)
	}
}
