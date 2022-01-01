package pkg

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DeleteChannelForCategory(s *discordgo.Session, category1 string, category2 string) {

	//fmt.Println("hello")

	for _, guild := range s.State.Guilds {

		// Get channels for this guild
		channels, _ := s.GuildChannels(guild.ID)

		var parentChannelId1 string
		var parentChannelId2 string
		for _, c1 := range channels {
			fmt.Println(c1.Type)
			// Check if channel is a guild text channel and not a voice or DM channel
			if strings.TrimRight(c1.Name, "\n") == category1 && c1.Type == 4 {

				parentChannelId1 = c1.ID

			} else if strings.TrimRight(c1.Name, "\n") == category2 && c1.Type == 4 {

				parentChannelId2 = c1.ID

			}
		}

		// Deleting the channels here

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if strings.TrimRight(c.ParentID, "\n") == parentChannelId1 && c.Type == 0 {
				s.ChannelDelete(c.ID)

			} else if strings.TrimRight(c.ParentID, "\n") == parentChannelId2 && c.Type == 0 {
				s.ChannelDelete(c.ID)

			}
		}
	}

}
