package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	// "OTIzODMzMjA2MzYzNTM3NDI5.YcVwuA.QOyVM2afeV-nGQPGGNsGlQa1ShA"
	Token string
)

func init() {
	godotenv.Load()
	Token = os.Getenv("token")
	fmt.Println(Token)
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, guild := range s.State.Guilds {

		// Get channels for this guild
		channels, _ := s.GuildChannels(guild.ID)

		var parentChannelId1 string
		var parentChannelId2 string
		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if strings.TrimRight(c.Name, "\n") == "tickets" {
				parentChannelId1 = c.ID

			} else if strings.TrimRight(c.Name, "\n") == "tickets2" {
				parentChannelId2 = c.ID

			}

		}

		// Deleting the channels here

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if strings.TrimRight(c.ParentID, "\n") == parentChannelId1 {
				s.ChannelDelete(c.ID)

			} else if strings.TrimRight(c.ParentID, "\n") == parentChannelId2 {
				s.ChannelDelete(c.ID)

			}

		}
	}

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {

		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
