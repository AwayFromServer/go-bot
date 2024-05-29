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
	BotToken string
)

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
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
	if message.Author.ID == discord.State.User.ID {
		return
	}
	switch {
	case strings.Contains(message.ChannelID, "1203220338100535327"):
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
			currentStatus := getCurrentStatus()
			_, err := discord.ChannelMessageSendComplex(message.ChannelID, currentStatus)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
