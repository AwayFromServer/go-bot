package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/awayfromserver/gobot/bot"
)

func main() {
	err := getBotToken()
	if err != nil {
		log.Fatal(err)
	}
	err = getTargetUrl()
	if err != nil {
	}
}

func getBotToken() error {
	// set BOT_TOKEN
	bt, ok := os.LookupEnv("BOT_TOKEN")
	if !ok || bt == "" {
		return fmt.Errorf("must set %s as env variable", "discord token")
	}
	return nil
}
func getTargetUrl(bt string) error {
	// set TARGET_URL
	t, ok := os.LookupEnv("TARGET_URL")
	if !ok || t == "" {
		return fmt.Errorf("must set %s as env variable", "discord token")
	}

	return run(bt, t)
}

func run(bt, t string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b := bot.New("", "")
	err := b.Run(ctx)

	botChannel := make(chan os.Signal, 1)
	signal.Notify(botChannel, os.Interrupt)
	<-botChannel
	fmt.Println("Bot shutting down...")

	return err
}
