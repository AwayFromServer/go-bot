package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/awayfromserver/gobot/bot"
)

var (
	BotChannel = make(chan os.Signal, 1)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Bot setup...")
	bot.Setup()
	fmt.Println("Bot starting up...")
	bot.Run(ctx)

	signal.Notify(BotChannel, os.Interrupt)
	<-BotChannel
	fmt.Println("Bot shutting down...")
	cancel()
}
