package commands

import (
	"errors"
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/errs"
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

	if betAmount <= 0 {
		s.ChannelMessageSend(m.ChannelID, "You must bet at least $1!")
		return
	}

	won, err := gambling.FlipCoin(player, betAmount)

	if errors.Is(err, errs.NoMoney) {
		s.ChannelMessageSend(m.ChannelID, "You don't have enough money! If you're out of money, you can declare bankruptcy.")
		return
	}

	if won {
		s.ChannelMessageSend(m.ChannelID, "Heads, You won $"+util.ToString(betAmount)+" losing "+util.ToString(casino.GetLoanPercentage(player))+"% to loan interest. You now have $"+util.ToString(player.Balance))
	} else {
		s.ChannelMessageSend(m.ChannelID, "Tails, You lost $"+util.ToString(betAmount)+". You now have $"+util.ToString(player.Balance))
	}
}
