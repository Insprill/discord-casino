package casino

type Player struct {
	ID      string
	Balance int64
}

//CheckBalance Checks if the players balance is at least the provided amount.
func CheckBalance(player Player, amount int64) bool {
	return player.Balance >= amount
}
