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

	if len(args) < 3 {
		s.ChannelMessageSend(m.ChannelID, "You must provide a target and an amount to remove!")
		return
	}

	if len(m.Mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, "You must provide a target by mentioning them!")
		return
	}

	target := m.Mentions[0]

	amount, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid amount "+args[1])
		return
	}

	player := casino.GetPlayer(target)
	player.Balance += amount

	s.ChannelMessageSend(m.ChannelID, "Successfully removed $"+util.ToString(amount)+" from "+target.Username+".")
}
