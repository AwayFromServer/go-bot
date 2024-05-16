package bot

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

const URL string = "http://hadbrew.my.pebble.host:8012"

func getCurrentStatus(message string) *discordgo.MessageSend {
	client := http.Client{Timeout: 5 * time.Second}

	response, err := client.Get(URL)
	if response == nil || err != nil {
		return &discordgo.MessageSend{
			Content: "I dunno about the server, but Dynmap appears to be offline... " + URL,
		}
	}

	return &discordgo.MessageSend{
		Content: "It looks like it's up! " + URL,
	}
}
