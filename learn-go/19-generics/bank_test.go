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

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := generics.Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}
