package commands

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/util"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func removeMoney(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	perms, err := s.State.MessagePermissions(m.Message)
	if err != nil || perms&discordgo.PermissionAdministrator == 0 {
		s.ChannelMessageSend(m.ChannelID, "You don't have permission to do that!")
		return
	}

	amount, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid amount "+args[1])
		return
	}

	player := casino.GetPlayer(m.Author)
	player.Balance -= amount
	if player.Balance < 0 {
		player.Balance = 0
	}

	s.ChannelMessageSend(m.ChannelID, "You now have $"+util.ToString(player.Balance))
}
