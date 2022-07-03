package casino

import (
	"github.com/Insprill/discord-casino/errs"
	"math"
)

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

func TakeLoan(player *Player, amount int64) {
	player.Balance += amount
	player.Loan += amount
	IncreaseLoanInterest(player, 0.05)
}

func RepayLoan(player *Player, amount int64) error {
	if !CheckBalance(player, amount) {
		return errs.NoMoney
	}
	player.Balance -= amount
	player.Loan -= amount
	if player.Loan <= 0 {
		player.Loan = 0
		player.LoanInterest = 1
	}
	return nil
}

func AddMoneyFromWin(player *Player, amount int64) {
	player.Balance += int64(float64(amount) * player.LoanInterest)
}

func RemoveMoneyFromLoss(player *Player, amount int64) {
	player.Balance -= amount
	IncreaseLoanInterest(player, 0.10)
}

func IncreaseLoanInterest(player *Player, amount float64) {
	player.LoanInterest = math.Max(0, math.Min(player.LoanInterest-amount, 0.95))
}

func GetLoanPercentage(player *Player) int64 {
	return int64(math.Abs((1 - player.LoanInterest) * 100))
}

func DeclareBankruptcy(player *Player) {
	player.Balance = 0
	player.LoanInterest = 1
	TakeLoan(player, 50)
}

func CreatePlayer(userId string) *Player {
	player := &Player{
		ID: userId,
	}
	DeclareBankruptcy(player)
	return player
}
