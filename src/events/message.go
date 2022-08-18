package events

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/config"
	"github.com/anhgelus/discord-hogwarts-housing/src/events/message/commands"
	"github.com/anhgelus/discord-hogwarts-housing/src/util"
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
	msg := fmt.Sprintf(m.Author.Username+" given xp %v", xp.NewMessage(m.Message))
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		panic(err)
	}
}

func prefixied(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	if util.StartWith(content, "ping") {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			panic(err)
		}
		return
	} else if util.StartWith(content, "help") {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, generateHelpEmbed())
		if err != nil {
			panic(err)
		}
	} else if util.StartWith(content, "house") {
		commands.HouseCommand(s, m)
	}
}

func generateHelpEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Help Page",
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "ping",
				Value: "Send the ping of the bot",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "help",
				Value: "Show this message",
			},
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Hogwarts Housing",
		},
	}
}
