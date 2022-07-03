package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
)

func userInfo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	player := casino.GetPlayer(m.Author)

	balanceField := &discordgo.MessageEmbedField{
		Name:   "Balance",
		Value:  "$" + util.ToString(player.Balance),
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

	embed := &discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{balanceField, loanSizeField, loanInterestField},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
