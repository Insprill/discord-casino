package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func takeLoan(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please provide how large of a loan you want to take out!")
		return
	}

	amount, err := strconv.ParseInt(args[1], 10, 64)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, args[1]+" is not a valid amount!")
		return
	}

	player := casino.GetPlayer(m.Author)

	casino.TakeLoan(player, amount)

	s.ChannelMessageSend(m.ChannelID, "Successfully took out a $"+util.ToString(amount)+" loan! You now have $"+util.ToString(player.Balance))
}
