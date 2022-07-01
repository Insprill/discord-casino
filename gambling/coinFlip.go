package gambling

import (
	"github.com/Insprill/discord-casino/casino"
	"github.com/Insprill/discord-casino/errs"
	"math/rand"
	"time"
)

// FlipCoin Flips a coin and either adds or removes the bet amount.
// Returns true if the player won, false otherwise
func FlipCoin(player *casino.Player, betAmount int64) (bool, error) {
	if !casino.CheckBalance(player, betAmount) {
		return false, errs.NoMoney
	}
	rand.Seed(time.Now().UnixNano())
	choseHeads := rand.Intn(2) == 0
	if choseHeads {
		casino.AddMoneyFromWin(player, betAmount)
		return true, nil
	} else {
		casino.RemoveMoneyFromLoss(player, betAmount)
		return false, nil
	}
}
