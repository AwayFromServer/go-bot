package bot

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func getCurrentStatus(url string) (*discordgo.MessageSend, error) {
	client := http.Client{Timeout: 5 * time.Second}
	content := ""

	response, err := client.Get(url)
	if response == nil || err != nil {
		content = "It looks like it's offline..."
	} else {
		content = "It looks like it's up!"
	}

	return &discordgo.MessageSend{
		Content: content,
	}, err
}
