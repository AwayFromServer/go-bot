package main

import (
	"go-bot/bot"
	"log"
	"os"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	bot.BotToken = botToken
	bot.Run()
}
