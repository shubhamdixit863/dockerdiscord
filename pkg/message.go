package pkg

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load()

}
//Will Reply to pong for Ping

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "deleteIt" {
		//Deleting the channels here
		DeleteChannelForCategory(s, os.Getenv("category1"), os.Getenv("category2"))

		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
