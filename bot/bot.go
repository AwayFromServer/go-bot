package bot

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	botToken  string
	targetURL string
}

func New(bt, t string) *Bot {
	b := Bot{
		botToken:  bt,
		targetURL: t,
	}
	return &b
}

func (b *Bot) Run(ctx context.Context) error {
	discord, err := discordgo.New("Bot " + b.botToken)
	if err != nil {
		return err
	}

	discord.AddHandler(b.newMessage)

	err = discord.Open()
	if err != nil {
		return err
	}
	defer discord.Close()

	fmt.Println("Bot running...")

	return err
}

func (b *Bot) newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Don't respond to bot's own messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

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
		currentStatus, terr := getCurrentStatus(b.targetURL)
		if terr != nil {
			log.Fatal(terr)
		}
		_, err := discord.ChannelMessageSendComplex(message.ChannelID, currentStatus)
		if err != nil {
			log.Fatal(err)
		}
	}
}
