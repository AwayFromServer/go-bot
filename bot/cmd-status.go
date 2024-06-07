package bot

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func getCurrentStatus(hbUrl string) (*discordgo.MessageSend, error) {
	client := http.Client{Timeout: 5 * time.Second}
	content := ""

	response, err := client.Get(hbUrl)
	if response == nil || err != nil {
		content = "It looks like it's offline... " + hbUrl
	} else {
		content = "It looks like it's up! " + hbUrl
	}

	return &discordgo.MessageSend{
		Content: content,
	}, err
}
