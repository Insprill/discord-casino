package gambling

import (
	"errors"
	"github.com/Insprill/discord-casino/casino"
	"math/rand"
)

// FlipCoin Flips a coin and either adds or removes the bet amount.
// Returns true if the player won, false otherwise
func FlipCoin(player *casino.Player, betAmount int64) (bool, error) {
	if !casino.CheckBalance(*player, betAmount) {
		return false, errors.New("not enough money")
	}
	choseHeads := rand.Intn(2) == 0
	if choseHeads {
		player.Balance += betAmount
		return true, nil
	} else {
		player.Balance -= betAmount
		return false, nil
	}
}
