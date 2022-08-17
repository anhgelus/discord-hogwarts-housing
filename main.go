package main

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var token = os.Args[1]

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	err = discord.UpdateGameStatus(0, "Je mange des jeans")
	if err != nil {
		panic(err)
	}

	src.RegisterHandler(discord)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = discord.Close()
	if err != nil {
		panic(err)
	}
}
