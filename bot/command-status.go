package bot

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func getCurrentStatus(hbUrl string) (*discordgo.MessageSend, error) {
	client := http.Client{Timeout: 5 * time.Second}

	response, err := client.Get(hbUrl)
	if response == nil || err != nil {
		return &discordgo.MessageSend{
			Content: "It looks like it's offline... " + hbUrl,
		}, nil
	}

	return &discordgo.MessageSend{
		Content: "It looks like it's up! " + hbUrl,
	}, nil
}
