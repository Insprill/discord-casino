package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Message.Content, "!casino") {
		return
	}

	args := strings.Split(m.Message.Content, " ")[1:]
	if len(args) == 0 {
		_, err := s.ChannelMessageSend(m.ChannelID, "Please provide a command to run!")
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	switch strings.ToLower(args[0]) {
	case "balance":
		{
			showBalance(s, m, args)
		}
	case "givemoney":
		{
			giveMoney(s, m, args)
		}

	case "removemoney":
		{
			removeMoney(s, m, args)
		}
	case "coinflip":
		{
			flipCoin(s, m, args)
		}
	}

}
