package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken  string
	targetURL string

	err error
)

func setup() error {
	bt, ok := os.LookupEnv("BOT_TOKEN")

	if !ok || bt == "" {
		err = fmt.Errorf("must set %a as env variable", "discord token")
	}
	tURL, ok := os.LookupEnv("TARGET_URL")
	if !ok || tURL == "" {
		err = fmt.Errorf("must set %a as env variable", "discord token")
	}
	log.Print("Setting botToken and targetURL")
	botToken = bt
	targetURL = tURL
	return err
}

func Run(ctx context.Context) error {
	setup()
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

	return err
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) error {
	// Don't respond to bot's own messages
	if message.Author.ID == discord.State.User.ID {
		return nil
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
			currentStatus, hberr := getCurrentStatus(targetURL)
			if hberr != nil {
				log.Fatal(hberr)
			}
			_, err := discord.ChannelMessageSendComplex(message.ChannelID, currentStatus)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return err
}
