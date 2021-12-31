package pkg

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DeleteChannelForCategory(s *discordgo.Session, category1 string, category2 string) {

	for _, guild := range s.State.Guilds {

		// Get channels for this guild
		channels, _ := s.GuildChannels(guild.ID)

		var parentChannelId1 string
		var parentChannelId2 string
		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if strings.TrimRight(c.Name, "\n") == category1 {
				parentChannelId1 = c.ID

			} else if strings.TrimRight(c.Name, "\n") == category2 {
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

}
