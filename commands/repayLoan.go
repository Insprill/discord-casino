package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/status"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func repayLoan(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	player := casino.GetPlayer(m.Author)

	if player.Loan <= 0 {
		s.ChannelMessageSend(m.ChannelID, "You don't have a loan to pay back!")
		return
	}

	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Please provide how much of your loan you want to pay back!")
		return
	}

	amount, err := strconv.ParseInt(args[1], 10, 64)
	if amount > player.Loan {
		amount = player.Loan
	}

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, args[1]+" is not a valid amount!")
		return
	}

	stat := casino.RepayLoan(player, amount)
	if stat == status.NoMoney {
		s.ChannelMessageSend(m.ChannelID, "You don't have enough money to do that!")
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Successfully payed $"+util.ToString(amount)+" off your loan! You now have $"+util.ToString(player.Balance)+" and owe another $"+util.ToString(player.Loan)+".")
}
