package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
)

func loanInfo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	player := casino.GetPlayer(m.Author)
	s.ChannelMessageSend(m.ChannelID, "You have a loan of $"+util.ToString(player.Loan)+" with an interest rate of "+util.ToString(casino.GetLoanPercentage(player))+"%")
}
