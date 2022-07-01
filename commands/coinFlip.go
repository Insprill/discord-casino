package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/gambling"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func flipCoin(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please provide how much you want to bet!")
		return
	}

	player := casino.GetPlayer(m.Author)
	betAmount, err := strconv.ParseInt(args[1], 10, 64)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, args[1]+" is not a valid amount!")
		return
	}

	won, err := gambling.FlipCoin(player, betAmount)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "You don't have enough money!")
		return
	}

	if won {
		s.ChannelMessageSend(m.ChannelID, "Heads, You won $"+util.ToString(betAmount)+"!")
	} else {
		s.ChannelMessageSend(m.ChannelID, "Tails, You lost $"+util.ToString(betAmount)+".")
	}
}
