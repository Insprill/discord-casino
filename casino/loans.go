package casino

import (
	"github.com/Insprill/discord-casino/status"
	"math"
)

var (
	MaxLoan     int64 = 10_000
	MaxInterest       = 0.05
)

func TakeLoan(player *Player, amount int64) int8 {
	if player.Loan >= MaxLoan {
		return status.MaxLoan
	}
	player.Balance += amount
	player.Loan += amount
	TryIncreaseLoanInterest(player, 0.05)
	return status.Ok
}

func RepayLoan(player *Player, amount int64) int8 {
	if !CheckBalance(player, amount) {
		return status.NoMoney
	}
	player.Balance -= amount
	player.Loan -= amount
	if player.Loan <= 0 {
		player.Loan = 0
		player.LoanInterest = 1
	}
	return status.Ok
}

func TryIncreaseLoanInterest(player *Player, amount float64) {
	if player.Loan > 0 {
		player.LoanInterest = math.Max(MaxInterest, math.Min(player.LoanInterest-amount, 1))
	}
}

func GetLoanPercentage(player *Player) int64 {
	return int64(math.Abs((1 - player.LoanInterest) * 100))
}

func DeclareBankruptcy(player *Player) {
	player.Balance = 0
	player.Loan = 0
	player.LoanInterest = 1
	TakeLoan(player, 50)
}
