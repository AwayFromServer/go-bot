package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/awayfromserver/gobot/bot"
)

var (
	BotChannel = make(chan os.Signal, 1)
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func getFromEnv(target string) (string, error) {
	value, ok := os.LookupEnv(target)
	if !ok || value == "" {
		return "", errors.New("must set " + target + " as env variable")
	}
	return value, nil
}

func run() error {
	bt, err := getFromEnv("BOT_TOKEN")
	if err != nil {
		return err
	}

	t, err := getFromEnv("TARGET_URL")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b := bot.New(bt, t)
	go b.Run(ctx)

	signal.Notify(BotChannel, os.Interrupt)
	<-BotChannel
	fmt.Println("Bot shutting down...")

	return err
}
