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
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Run(ctx)

	signal.Notify(BotChannel, os.Interrupt)
	<-BotChannel
	fmt.Println("Bot shutting down...")

	return err
}
