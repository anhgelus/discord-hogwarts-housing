package xp

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/database"
	"github.com/anhgelus/discord-hogwarts-housing/src/util"
	"github.com/bwmarrin/discordgo"
	"math"
)

func NewMessage(m *discordgo.Message) float64 {
	c := database.GetRedisPool()
	if c == nil {
		return 0
	}
	key := "xp:" + m.GuildID + ":" + m.Author.ID
	xp := calc(len(m.Content), util.GetNumberOfChar(m.Content))
	_, err := database.RedisSet(c, key, fmt.Sprintf("%f", xp))
	if err != nil {
		return 0
	}
	return xp
}

// l int - length of the message
// v int - Number of character in the message
func calc(l int, v int) float64 {
	// f(x)=((0.025 x^(1.25))/(50^(-0.5)))+1
	result := 0.025 * math.Pow(float64(l), 1.25)
	result = result / math.Pow(float64(v), -0.5)
	return result + 1
}
