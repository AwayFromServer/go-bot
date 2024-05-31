package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken     string
	heartbeatURL string
)

func Setup() {
	bt, ok := os.LookupEnv("BOT_TOKEN")
	if !ok || bt == "" {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}
	hbURL, ok := os.LookupEnv("HEARTBEAT_URL")
	if !ok || hbURL == "" {
		log.Fatal("Must set heartbeat URL as env variable: HEARTBEAT_URL")
	}
	botToken = bt
	heartbeatURL = hbURL
}

func Run() {
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Don't respond to bot's own messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Outer switch looking for specific channels
	switch {
	case strings.Contains(message.ChannelID, "1203220338100535327"): // #gaming-stuff
		// Inner switch looking for specific commands issued as messages in the aforementioned channel
		switch {
		case strings.Contains(message.Content, "server status"):
			_, err := discord.ChannelMessageSend(message.ChannelID, "I can help with that! Use '!status'!")

			if err != nil {
				log.Fatal(err)
			}
		case strings.Contains(message.Content, "bot"):
			_, err := discord.ChannelMessageSend(message.ChannelID, "Who, me?")
			if err != nil {
				log.Fatal(err)
			}
		case strings.Contains(message.Content, "!status"):
			currentStatus, hberr := getCurrentStatus(heartbeatURL)
			if hberr != nil {
				log.Fatal(hberr)
			}
			_, err := discord.ChannelMessageSendComplex(message.ChannelID, currentStatus)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
