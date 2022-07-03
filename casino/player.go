package casino

type Player struct {
	ID           string
	Balance      int64
	Loan         int64
	LoanInterest float64
}

//CheckBalance Checks if the players balance is at least the provided amount.
func CheckBalance(player *Player, amount int64) bool {
	return player.Balance >= amount
}

func AddMoneyFromWin(player *Player, amount int64) {
	player.Balance += int64(float64(amount) * player.LoanInterest)
}

func RemoveMoneyFromLoss(player *Player, amount int64) {
	player.Balance -= amount
	TryIncreaseLoanInterest(player, 0.10)
}

func CreatePlayer(userId string) *Player {
	player := &Player{
		ID: userId,
	}
	DeclareBankruptcy(player)
	return player
}
