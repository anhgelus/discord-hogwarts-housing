package src

import (
	"github.com/anhgelus/discord-hogwarts-housing/src/events"
	"github.com/bwmarrin/discordgo"
)

func RegisterHandler(discord *discordgo.Session) {
	discord.AddHandler(events.MessageCreate)
}
