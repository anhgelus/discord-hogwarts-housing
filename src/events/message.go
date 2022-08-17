package events

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/config"
	"github.com/anhgelus/discord-hogwarts-housing/src/xp"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, config.Prefix) {
		prefixied(s, m)
		return
	}
	msg := fmt.Sprintf(m.Author.Username+" given xp %v", xp.NewMessage(m.Content))
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		panic(err)
	}
}

func prefixied(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	if startWith(content, "ping") {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			panic(err)
		}
		return
	} else if startWith(content, "help") {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "Hogwarts Housing",
			Color: 0x00ff00,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:  "ping",
					Value: "Send the ping of the bot",
				},
				&discordgo.MessageEmbedField{
					Name:  "help",
					Value: "Show this message",
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name: "Hogwarts Housing",
			},
		})
		if err != nil {
			panic(err)
		}
	}
}

func startWith(m string, c string) bool {
	return strings.HasPrefix(m, config.Prefix+c)
}
