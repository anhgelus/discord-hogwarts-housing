package commands

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/config"
	"github.com/anhgelus/discord-hogwarts-housing/src/guild"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HouseCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	content = strings.TrimPrefix(content, config.Prefix)
	args := strings.Split(content, " ")
	length := len(args)
	if length == 1 {
		helpHouse(s, m.ChannelID)
	} else if length == 2 {
		if args[1] == "list" {
			listHouse(s, m)
			return
		} else if args[1] == "leave" {
			// do a command
			return
		}
		helpHouse(s, m.ChannelID)
	} else if length == 3 {
		if args[1] == "create" {
			createHouse(s, m, args[2])
			return
		}
	}
}

func createHouse(s *discordgo.Session, m *discordgo.MessageCreate, name string) {
	house, err := guild.CreateHouse(m.GuildID, name, m.Author.ID)
	if err != nil {
		_, err2 := s.ChannelMessageSend(m.ChannelID, "An error occurred: `"+err.Error()+"`")
		if err2 != nil {
			panic(err2)
		}
		return
	}
	msg := fmt.Sprintf("House %s created with id %s", house.Name, house.Id)
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
}

func listHouse(s *discordgo.Session, m *discordgo.MessageCreate) {
	houses, err := guild.GetHouses(m.GuildID)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "There are no houses")
		if err != nil {
			panic(err)
		}
		return
	}
	for _, house := range houses {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` - **%s**", house.Id, house.Name))
		if err != nil {
			panic(err)
		}
	}
}

func leaveHouse(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	house, err := guild.GetHouse(args[2], m.GuildID)
	if err != nil {
		_, err2 := s.ChannelMessageSend(m.ChannelID, "An error occurred: `"+err.Error()+"`")
		if err2 != nil {
			panic(err2)
		}
	}
	err = guild.UserLeave(house.Id, m.Author.ID, m.GuildID)
	if err != nil {
		_, err2 := s.ChannelMessageSend(m.ChannelID, "An error occurred: `"+err.Error()+"`")
		if err2 != nil {
			panic(err2)
		}
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
				Name:  config.Prefix + "house create <name>", // done
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
				Name:  config.Prefix + "house list", // done
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
			&discordgo.MessageEmbedField{
				Name:  config.Prefix + "house exist <id>",
				Value: "Check if the house with the id <id> exists",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
