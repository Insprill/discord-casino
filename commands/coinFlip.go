package commands

import (
	"errors"
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/errs"
	"github.com/Insprill/discord-casino/gambling"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
	"math"
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

	balBefore := player.Balance

	won, err := gambling.FlipCoin(player, betAmount)

	if errors.Is(err, errs.NoMoney) {
		s.ChannelMessageSend(m.ChannelID, "You don't have enough money! If you're out of money, you can declare bankruptcy.")
		return
	}

	amountChanged := int64(math.Abs(float64(player.Balance - balBefore)))
	lostToLoan := &discordgo.MessageEmbedField{
		Name:   "Amount Lost to Loan",
		Value:  "$" + util.ToString(int64(math.Abs(float64(betAmount-amountChanged)))),
		Inline: false,
	}
	loanSizeField := &discordgo.MessageEmbedField{
		Name:   "Loan Amount",
		Value:  "$" + util.ToString(player.Loan),
		Inline: false,
	}
	loanInterestField := &discordgo.MessageEmbedField{
		Name:   "Loan Interest",
		Value:  util.ToString(casino.GetLoanPercentage(player)) + "%",
		Inline: false,
	}

	if won {
		embed := &discordgo.MessageEmbed{
			Title:       "You won!",
			Description: "You won $" + util.ToString(amountChanged) + ".",
			Fields:      []*discordgo.MessageEmbedField{loanSizeField, loanInterestField, lostToLoan},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	} else {
		embed := &discordgo.MessageEmbed{
			Title:       "You lost.",
			Description: "You lost $" + util.ToString(amountChanged) + ".",
			Fields:      []*discordgo.MessageEmbedField{loanSizeField, loanInterestField},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
