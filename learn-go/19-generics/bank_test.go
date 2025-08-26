package generics_test

import (
	"testing"

	"github.com/lin-br/go-lin/generics"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = generics.Account{Name: "Riya", Balance: 100}
		chris = generics.Account{Name: "Chris", Balance: 75}
		adil  = generics.Account{Name: "Adil", Balance: 200}

		transactions = []generics.Transaction{
			generics.NewTransaction(chris, riya, 100),
			generics.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account generics.Account) float64 {
		return generics.NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}
