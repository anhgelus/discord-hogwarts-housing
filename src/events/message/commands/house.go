package commands

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HouseCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	content = strings.TrimPrefix(content, config.Prefix)
	args := strings.Split(content, " ")
	length := len(args)
	for i, arg := range args {
		fmt.Printf("%d:%s\n", i, arg)
	}
	if length == 1 {
		helpHouse(s, m.ChannelID)
	} else if length == 2 {
		fmt.Printf("%s\n", args[1])
		if args[1] == "list" {
			// do a command
			return
		} else if args[1] == "leave" {
			// do a command
			return
		}
		helpHouse(s, m.ChannelID)
	}
}

// s = session
// id = channelID
func helpHouse(s *discordgo.Session, id string) {
	_, err := s.ChannelMessageSendEmbed(id, &discordgo.MessageEmbed{
		Title: "Help House",
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house create <name>",
				Value: "Create a new house, with the name <name>",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house delete <id>",
				Value: "Delete the house with the id <id>",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house update <id> <value> <new-value>",
				Value: "Update the house with the id <id> and change the value <value> by the new value <new-value>",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house list",
				Value: "List all houses",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house info <id>",
				Value: "Show the info of the house with the id <id>",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house join <id>",
				Value: "Join the house with the id <id>",
			},
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house leave",
				Value: "Leave your current house",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
