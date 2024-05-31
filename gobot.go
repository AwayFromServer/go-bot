package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/awayfromserver/gobot/bot"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	bot.Setup()
	bot.Run(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Bot shutting down...")
	cancel()
}
