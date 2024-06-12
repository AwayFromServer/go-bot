package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/awayfromserver/gobot/config"
)

const BT = "botToken"
const BP = "botPrefix"

type Bot struct {
	config     []map[string]interface{}
	session    *discordgo.Session
	botChannel chan os.Signal
}

/*
Run() is essentially the main bot loop
*/
func Run() error {
	cfgFilename := "config.json"

	for i, argument := range os.Args {
		switch {
		case argument == "-f":
			{
				log.Printf("Override detected for %s. New value: %s", "config file", os.Args[i+1])
				cfgFilename = os.Args[i+1]
			}
		}
	}

	cfg, err := config.GetConfig(cfgFilename)
	if err != nil {
		return err
	}
	log.Println("Finished reading the file!")

	b := Bot{config: config.GetOverrides(cfg)}
	log.Println("Finished reading overrides!")

	b.session, err = discordgo.New("Bot " + b.config[0][BT].(string))
	if err != nil {
		return err
	}
	log.Println("Session created!")

	b.session.AddHandler(newMessage)
	err = b.session.Open()
	log.Println("Session opened!")
	defer b.session.Close()
	if err != nil {
		return err
	}

	b.botChannel = make(chan os.Signal, 1)
	log.Println("Bot is now running...")
	log.Println("Ctrl+C to exit...")
	signal.Notify(b.botChannel, os.Interrupt)
	<-b.botChannel
	fmt.Println("Bot shutting down...")
	return err
}

func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	var err error
	switch {
	case message.Author.ID == session.State.User.ID:
		return // Don't respond to bot's own messages
	case strings.Contains(message.Content, "server status"):
		_, err = session.ChannelMessageSend(message.ChannelID, "I can help with that! Use '!status'!")
	case strings.Contains(message.Content, "bot"):
		_, err = session.ChannelMessageSend(message.ChannelID, "Who, me?")
	case string([]rune(message.Content)[0]) == "!":
		log.Printf("Command received: %s", message.Content)
		cmdLine := strings.Split(message.Content, "!")
		cmdWords := strings.Split(cmdLine[1], " ")

		err = execBotCommand(session, message, cmdWords)
	}
	if err != nil {
		log.Print(err)
	}
}

func execBotCommand(session *discordgo.Session, message *discordgo.MessageCreate, cmdWords []string) error {
	var err error
	var currentStatus *discordgo.MessageSend
	switch {
	case cmdWords[0] == "status":
		currentStatus, err = getCurrentStatus(cmdWords[1])
		if err == nil {
			_, err = session.ChannelMessageSendComplex(message.ChannelID, currentStatus)
		}
	}
	return err
}
