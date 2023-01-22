package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	Token     string
	ChannelId string
	Session   *discordgo.Session
}

func New(token string, c string) (*Discord, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return &Discord{
		Token:     token,
		ChannelId: c,
		Session:   dg,
	}, nil
}

func (d *Discord) Send(data string) {
	_, err := d.Session.ChannelMessageSend(d.ChannelId, data)
	if err != nil {
		log.Printf("Discord: failed to send message: %v", err.Error())
		return
	}
}
