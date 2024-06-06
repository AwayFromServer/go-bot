package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

const BT = "BOT_TOKEN"
const TU = "TARGET_URL"
const BP = "BOT_PREFIX"
const CFGFILE = "../config.yaml"

type Bot struct {
	config     conf
	session    *discordgo.Session
	botChannel chan os.Signal
}

type conf struct {
	BotToken  string `yaml:"Token"`
	BotTarget string `yaml:"Target"`
	BotPrefix string `yaml:"Prefix"`
}

func Run() {
	// read in config and overrides
	var c conf
	c.getConf(CFGFILE)
	c.getOverrides()

	// assign config to new Bot
	b := Bot{config: c}
	b.session = b.startSession()
	b.session.AddHandler(b.newMessage)
	// open session connection
	err := b.session.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer b.session.Close()

	b.botChannel = make(chan os.Signal, 1)
	signal.Notify(b.botChannel, os.Interrupt)
	<-b.botChannel
	fmt.Println("Bot shutting down...")
}

func (c *conf) getConf(filename string) *conf {

	yamlFile, err := os.ReadFile(filename)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, c)
	}

	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *conf) getOverrides() *conf {
	bt, btok := os.LookupEnv(BT)
	if !btok {
		log.Printf("No override detected for %s", BT)
	} else if bt != "" {
		log.Printf("Override detected for %s. Swapping for ENVVAR. New value: %s", BT, bt)
		c.BotToken = bt
	}

	tu, tuok := os.LookupEnv(TU)
	if !tuok {
		log.Printf("No override detected for %s", TU)
	} else if tu != "" {
		log.Printf("Override detected for %s. Swapping for ENVVAR. New value: %s", TU, tu)
		c.BotTarget = tu
	}

	bp, bpok := os.LookupEnv(BP)
	if !bpok {
		log.Printf("No override detected for %s", BP)
	} else if bp != "" {
		log.Printf("Override detected for %s. Swapping for ENVVAR. New value: %s", BP, bp)
		c.BotPrefix = bp
	}

	return c
}

func (b *Bot) startSession() *discordgo.Session {
	session, err := discordgo.New("Bot " + b.config.BotToken)

	if err != nil {
		log.Fatal(err)
	}

	return session
}

func (b *Bot) newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) error {
	var err error
	switch {
	case message.Author.ID == discord.State.User.ID:
		return nil // Don't respond to bot's own messages
	case strings.Contains(message.Content, "server status"):
		_, err = discord.ChannelMessageSend(message.ChannelID, "I can help with that! Use '!status'!")
	case strings.Contains(message.Content, "bot"):
		_, err = discord.ChannelMessageSend(message.ChannelID, "Who, me?")
	case strings.Contains(message.Content, "!status"):
		var currentStatus *discordgo.MessageSend
		currentStatus, err = getCurrentStatus(b.config.BotTarget)
		if err == nil {
			_, err = discord.ChannelMessageSendComplex(message.ChannelID, currentStatus)
		}
	}
	return err
}
