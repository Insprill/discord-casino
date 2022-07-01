package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
)

func declareBankruptcy(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	player := casino.GetPlayer(m.Author)
	casino.DeclareBankruptcy(player)
	s.ChannelMessageSend(m.ChannelID, "You filed for bankruptcy! You now only have $"+util.ToString(player.Balance)+" on a loan.")
}
