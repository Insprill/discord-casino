package casino

import (
	"github.com/Insprill/discord-casino/errs"
	"math"
)

func TakeLoan(player *Player, amount int64) {
	player.Balance += amount
	player.Loan += amount
	TryIncreaseLoanInterest(player, 0.05)
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

func TryIncreaseLoanInterest(player *Player, amount float64) {
	if player.Loan <= 0 {
		player.LoanInterest = math.Max(0, math.Min(player.LoanInterest-amount, 0.95))
	}
}

func GetLoanPercentage(player *Player) int64 {
	return int64(math.Abs((1 - player.LoanInterest) * 100))
}

func DeclareBankruptcy(player *Player) {
	player.Balance = 0
	player.LoanInterest = 1
	TakeLoan(player, 50)
}
