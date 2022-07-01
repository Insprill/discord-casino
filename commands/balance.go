package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
)

func showBalance(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "You have $"+util.ToString(casino.GetPlayer(m.Author).Balance))
}
